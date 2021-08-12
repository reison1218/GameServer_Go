package main

import (
	"activitycenter/redis_helper"
	"activitycenter/season"
	"activitycenter/template"
	"activitycenter/worldboss"
)

func main() {
	//初始化配置表
	initTemplates()
	//初始化redis
	initReids()
	//初始化赛季
	initSeason()
	//初始化worldboss
	initWorldBoss()
	select {}
}

func initTemplates() {
	var a, b = template.NewSeasonTemplateMgr(), template.NewWorldBossTemplateMgr()
	template.TemplateGlobalMgr = template.TemplatesMgr{SeasonMgr: a, WorldBossMgr: b}
	tempMgr := template.TemplateGlobalMgr
	tempMgr.Init()
}

func initReids() {
	redis_helper.RedisGlobalHelper = redis_helper.Init()
}

func initSeason() {
	season.Init()
}

func initWorldBoss() {
	worldboss.Init()
}
