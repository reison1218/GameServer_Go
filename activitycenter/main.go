package main

import (
	"activitycenter/template"
)

func main() {
	initTemplates()
}

func initTemplates() {
	var a, b = template.NewSessionTemplateMgr(), template.NewWorldBossTemplateMgr()
	var tempMgr = template.TemplatesMgr{SessionMgr: a, WorldBossMgr: b}
	tempMgr.Init()
}
