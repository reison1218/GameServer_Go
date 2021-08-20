package main

import (
	"activitycenter/config_helper"
	"activitycenter/redis_helper"
	"activitycenter/season"
	"activitycenter/template_mgr"
	"activitycenter/worldboss"
	"log"
	"time"
)

func main() {
	startTime := time.Now()
	//初始化配置文件
	initConfig()
	//初始化配置表
	initTemplates()
	//初始化redis
	initReids()
	//初始化赛季
	initSeason()
	//初始化worldboss
	initWorldBoss()
	endTime := time.Now()
	res := endTime.Sub(startTime)
	log.Println("activitycenter init success!take time:", res)
	select {}
}

func initTemplates() {
	var a, b = template_mgr.NewSeasonTemplateMgr(), template_mgr.NewWorldBossTemplateMgr()
	template_mgr.TemplateGlobalMgr = template_mgr.TemplatesMgr{SeasonMgr: a, WorldBossMgr: b}
	tempMgr := template_mgr.TemplateGlobalMgr
	tempMgr.Init()
}

func initConfig() {
	config_helper.Init()
}

func initReids() {
	redis_helper.Init()
}

func initSeason() {
	season.Init()
}

func initWorldBoss() {
	worldboss.Init()
}
