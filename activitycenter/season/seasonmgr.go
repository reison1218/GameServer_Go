package season

import (
	"activitycenter/redis_helper"
	"activitycenter/template"
	"time"

	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
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
	SeasonMap map[int]SeasonInfo
}

func Init() {
	templateMgr := template.TemplateGlobalMgr
	redisHelper := redis_helper.RedisGlobalHelper
	redisHelper.Do("select", 1)

	res, err := redis.Values(redisHelper.Do("hgetall", "game_season"))
	if err != nil {
		panic(err)
	}
	nowTime := time.Now().UTC()
	seasonInfo := SeasonInfo{}
	needUpdate := false
	//执行初始化
	if len(res) == 0 {
		seasonInfo.GameId = 101
		seasonInfo.Round = 1
		seasonInfo.SeasonId = 1001
		needUpdate = true
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
			template := templateMgr.SeasonMgr.GetById(nextSeasonId + 1)
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
		res, _ := time.ParseDuration("2160h")
		seasonInfo.NextUpdateTime = nowTime.Add(res).Unix()
		seasonInfo.NextUpdateTimeStr = nowTime.Add(res).String()
		jsonRes, err := jsoniter.Marshal(seasonInfo)
		if err != nil {
			panic(err)
		}
		redisHelper.Do("hset", "game_season", "101", string(jsonRes))
	}
	SeasonGlobalMgr = newSeasonMgr()
	SeasonGlobalMgr.SeasonMap[seasonInfo.GameId] = seasonInfo
}

func newSeasonMgr() SeasonMgr {
	return SeasonMgr{SeasonMap: make(map[int]SeasonInfo)}
}
