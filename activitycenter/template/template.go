package template

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

var TemplateGlobalMgr TemplatesMgr

type Template interface {
}

type TemplateMgr interface {
	GetById(id int) Template

	IsEmpty() bool

	Clear()
}
type TemplatesMgr struct {
	SeasonMgr    SeasonTemplateMgr
	WorldBossMgr WorldBossTemplateMgr
}

func (mgr *TemplatesMgr) Init() {
	tempDir, err := ioutil.ReadDir("/Users/tangjian/Desktop/json")
	if err != nil {
		println(err)
		return
	}

	for _, file := range tempDir {
		fileName := file.Name()
		path := "/Users/tangjian/Desktop/json/" + fileName
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			println(err)
			continue
		}

		if fileName == "Season.json" {
			var seasonTemps []SeasonTemplate
			err := jsoniter.Unmarshal(bytes, &seasonTemps)
			if err != nil {
				println(err)
			}
			mgr.SeasonMgr.init(seasonTemps)
		} else if fileName == "WorldBoss.json" {
			var worldBossTemps []WorldBossTemplate
			err := jsoniter.Unmarshal(bytes, &worldBossTemps)
			if err != nil {
				println(err)
			}
			mgr.WorldBossMgr.init(worldBossTemps)
		}
	}
}
