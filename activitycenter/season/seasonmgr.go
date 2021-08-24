package season

import (
	"activitycenter/config_helper"
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

var SeasonGlobalMgr SeasonMgr

type SeasonInfo struct {
	GameId            int    `json:"game_id"`
	SeasonId          int    `json:"season_id"`
	Round             int    `json:"round_id"`
	LastUpdateTime    int64  `json:"last_update_time"`
	LastUpdateTimeStr string `json:"last_update_time_str"`
	NextUpdateTime    int64  `json:"next_update_time"`
	NextUpdateTimeStr string `json:"next_update_time_str"`
}

type SeasonMgr struct {
	SeasonInfo SeasonInfo
}

func Init() {
	gameId := config_helper.Configuration.Configs["game_id"]
	templateMgr := template_mgr.TemplateGlobalMgr
	redisHelper := redis_helper.RedisGlobalHelper
	redisHelper.Do("select", 1)

	res, err := redis.Values(redisHelper.Do("hgetall", "game_season"))
	if err != nil {
		panic(err)
	}
	nowTime := time.Now().UTC()
	seasonInfo := SeasonInfo{}
	needUpdate := false
	var template *template_mgr.SeasonTemplate
	//执行初始化
	if len(res) == 0 {
		res, _ := gameId.Int64()
		seasonInfo.GameId = int(res)
		seasonInfo.Round = 1
		seasonInfo.SeasonId = 1001
		needUpdate = true
		template = templateMgr.SeasonMgr.GetById(1001)
	} else {
		//如果没有就判断是否过期了没
		redisData := res[1]
		bytes := redisData.([]byte)
		err := jsoniter.Unmarshal(bytes, &seasonInfo)
		if err != nil {
			panic(err)
		}
		if nowTime.Unix() >= seasonInfo.NextUpdateTime {
			nextSeasonId := seasonInfo.SeasonId
			template = templateMgr.SeasonMgr.GetById(nextSeasonId + 1)
			if template.Id == 0 {
				nextSeasonId = 1001
			}
			seasonInfo.SeasonId = nextSeasonId
			needUpdate = true
		}
	}

	if needUpdate {
		seasonInfo.LastUpdateTime = nowTime.Unix()
		seasonInfo.LastUpdateTimeStr = nowTime.String()
		endTime := nowTime.Add(time.Duration(template.KeepTime) * time.Millisecond)
		seasonInfo.NextUpdateTime = endTime.Unix()
		seasonInfo.NextUpdateTimeStr = endTime.String()
		jsonRes, err := jsoniter.Marshal(seasonInfo)
		if err != nil {
			panic(err)
		}
		redisHelper.Do("hset", "game_season", "101", string(jsonRes))
	}
	SeasonGlobalMgr = newSeasonMgr()
	SeasonGlobalMgr.SeasonInfo = seasonInfo
	timer := newWithSeconds()
	go func() {
		for {
			res := seasonInfo.NextUpdateTime - seasonInfo.LastUpdateTime
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
	log.Println("season init success!")
}

func newSeasonMgr() SeasonMgr {
	return SeasonMgr{}
}

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func check_update() {
	gameId := config_helper.Configuration.Configs["game_id"]
	nowTime := time.Now().UTC()
	templateMgr := template_mgr.TemplateGlobalMgr
	redisHelper := redis_helper.RedisGlobalHelper
	seasonInfo := SeasonGlobalMgr.SeasonInfo
	if nowTime.Unix() < SeasonGlobalMgr.SeasonInfo.NextUpdateTime {
		return
	}
	seasonId := seasonInfo.SeasonId
	worldBossTemplate := templateMgr.SeasonMgr.GetNext(seasonId)
	seasonInfo.SeasonId = worldBossTemplate.Id
	seasonInfo.LastUpdateTime = nowTime.Unix()
	seasonInfo.LastUpdateTimeStr = nowTime.String()
	keepTime := strconv.FormatInt(worldBossTemplate.KeepTime, 10)
	res, _ := time.ParseDuration(keepTime + "ms")
	seasonInfo.NextUpdateTime = nowTime.Add(res).Unix()
	seasonInfo.NextUpdateTimeStr = nowTime.Add(res).String()
	jsonRes, err := jsoniter.Marshal(&seasonInfo)
	if err != nil {
		panic(err)
	}
	redisHelper.Do("hset", "game_season", gameId.String(), string(jsonRes))
	httpUrl := config_helper.Configuration.Configs["game_center_http"]
	//通知游戏服务器worldboss更新
	httpRes, err := http.Get(httpUrl.String() + "update_season")
	if err != nil {
		println(err)
		return
	}
	if httpRes.StatusCode == 200 {
		println("notify game server update season success!")
	} else {
		println("notify game server update fail!")
	}
}
