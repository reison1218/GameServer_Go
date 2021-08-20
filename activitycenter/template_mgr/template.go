package template_mgr

import (
	"io/ioutil"
	"log"
	"os"

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

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	jsonDir := dir + "/template/"
	tempDir, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		println(err)
		return
	}

	for _, file := range tempDir {
		fileName := file.Name()
		path := jsonDir + fileName
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
	log.Println("template init success!")
}
