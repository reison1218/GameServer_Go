package main

import (
	"activitycenter/redis"
	"activitycenter/template"
)

func main() {
	//initTemplates()
	redis.RedisTest()
}

func initTemplates() {
	var a, b = template.NewSessionTemplateMgr(), template.NewWorldBossTemplateMgr()
	var tempMgr = template.TemplatesMgr{SessionMgr: a, WorldBossMgr: b}
	tempMgr.Init()
}
