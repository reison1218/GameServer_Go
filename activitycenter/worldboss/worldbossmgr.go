package worldboss

import (
	"activitycenter/redis_helper"
	"activitycenter/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
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
	templateMgr := template.TemplateGlobalMgr
	redisHelper := redis_helper.RedisGlobalHelper
	nowTime := time.Now().UTC()
	redisHelper.Do("select", 1)
	res, err := redis.Values(redisHelper.Do("hgetall", "world_boss"))
	if err != nil {
		panic(err)
	}
	worldBossInfo := WorldBossInfo{}
	needUpdate := false
	var worldBossTemplate template.WorldBossTemplate
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
		redisHelper.Do("hset", "world_boss", worldBossInfo.CterId, string(jsonRes))
	}
	WorldBossGlobalMgr.WorldBossInfo = worldBossInfo

	sleepTime, _ := time.ParseDuration("2000ms")
	time.AfterFunc(sleepTime, check_update)
}

func check_update() {
	nowTime := time.Now().UTC()
	templateMgr := template.TemplateGlobalMgr
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
	redisHelper.Do("hset", "world_boss", worldBossInfo.CterId, string(jsonRes))
	//通知游戏服务器worldboss更新
	http.Get("http://127.0.0.1:9999/update_world_boss")
}
