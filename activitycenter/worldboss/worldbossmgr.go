package worldboss

import (
	"activitycenter/redis_helper"
	"activitycenter/template_mgr"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/robfig/cron/v3"
)

var WorldBossGlobalMgr WorldBossMgr

type WorldBossInfo struct {
	CterId            int    `json:"cter_id"`
	LastUpdateTime    int64  `json:"last_update_time"`
	LastUpdateTimeStr string `json:"last_update_time_str"`
	NextUpdateTime    int64  `json:"next_update_time"`
	NextUpdateTimeStr string `json:"next_update_time_str"`
}

type WorldBossMgr struct {
	WorldBossInfo WorldBossInfo
}

func newWorldBossMgr() WorldBossMgr {
	return WorldBossMgr{}
}

func Init() {
	WorldBossGlobalMgr = newWorldBossMgr()
	//先加载reids
	templateMgr := template_mgr.TemplateGlobalMgr
	redisHelper := redis_helper.RedisGlobalHelper
	nowTime := time.Now().UTC()
	redisHelper.Do("select", 1)
	res, err := redis.Values(redisHelper.Do("hgetall", "world_boss"))
	if err != nil {
		panic(err)
	}
	worldBossInfo := WorldBossInfo{}
	needUpdate := false
	var worldBossTemplate template_mgr.WorldBossTemplate
	if len(res) == 0 {
		worldBossTemplate = templateMgr.WorldBossMgr.GetFirst()
		worldBossInfo.CterId = worldBossTemplate.CterId
		needUpdate = true
	} else {
		bytes := res[1]
		err := jsoniter.Unmarshal(bytes.([]byte), &worldBossInfo)
		if err != nil {
			panic(err)
		}
		cterId := worldBossInfo.CterId
		//判断过期了没
		if nowTime.Unix() >= worldBossInfo.NextUpdateTime {
			needUpdate = true
			worldBossTemplate = templateMgr.WorldBossMgr.GetNext(cterId)
			worldBossInfo.CterId = worldBossTemplate.CterId
		}
	}
	if needUpdate {
		worldBossInfo.LastUpdateTime = nowTime.Unix()
		worldBossInfo.LastUpdateTimeStr = nowTime.String()
		keepTime := strconv.FormatInt(worldBossTemplate.KeepTime, 10)
		res, _ := time.ParseDuration(keepTime + "ms")
		worldBossInfo.NextUpdateTime = nowTime.Add(res).Unix()
		worldBossInfo.NextUpdateTimeStr = nowTime.Add(res).String()

		jsonRes, err := jsoniter.Marshal(&worldBossInfo)
		if err != nil {
			panic(err)
		}
		redisHelper.Do("hset", "world_boss", "101", string(jsonRes))
	}

	WorldBossGlobalMgr.WorldBossInfo = worldBossInfo
	timer := newWithSeconds()
	go func() {
		for {
			res := worldBossInfo.NextUpdateTime - worldBossInfo.LastUpdateTime
			sleepTime := strconv.FormatInt(res, 10)
			spec := "*/" + sleepTime + " * * * * ?"
			id, _ := timer.AddFunc(spec, check_update)
			timer.Start()
			//让协程等待200ms让任务执行完
			res += 200
			time.Sleep(time.Duration(res) * time.Millisecond)
			//删掉任务，从新执行
			timer.Remove(id)
			continue
		}
	}()
	log.Println("world_boss init success!")
}

func check_update() {
	nowTime := time.Now().UTC()
	templateMgr := template_mgr.TemplateGlobalMgr
	redisHelper := redis_helper.RedisGlobalHelper
	worldBossInfo := WorldBossGlobalMgr.WorldBossInfo
	if nowTime.Unix() < WorldBossGlobalMgr.WorldBossInfo.NextUpdateTime {
		return
	}
	cterId := worldBossInfo.CterId
	worldBossTemplate := templateMgr.WorldBossMgr.GetNext(cterId)
	worldBossInfo.CterId = worldBossTemplate.CterId
	worldBossInfo.LastUpdateTime = nowTime.Unix()
	worldBossInfo.LastUpdateTimeStr = nowTime.String()
	keepTime := strconv.FormatInt(worldBossTemplate.KeepTime, 10)
	res, _ := time.ParseDuration(keepTime + "ms")
	worldBossInfo.NextUpdateTime = nowTime.Add(res).Unix()
	worldBossInfo.NextUpdateTimeStr = nowTime.Add(res).String()
	jsonRes, err := jsoniter.Marshal(&worldBossInfo)
	if err != nil {
		panic(err)
	}
	redisHelper.Do("hset", "world_boss", "101", string(jsonRes))
	//通知游戏服务器worldboss更新
	http.Get("http://127.0.0.1:9999/update_world_boss")
}

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
