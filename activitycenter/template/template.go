package template

import (
	"time"
)

type Template interface {
}

type TemplateMgr interface {
	getById(id uint32) Template

	isEmpty() bool

	clear()
}
type TemplatesMgr struct {
	SessionMgr   SessionTemplateMgr
	WorldBossMgr WorldBossTemplateMgr
}

func (mgr TemplatesMgr) Init() {
	// tempDir, err := ioutil.ReadDir("/Users/tangjian/Desktop/test")
	// if err != nil {
	// 	println(err)
	// 	return
	// }

	// for _, file := range jsonDir {
	// 	path := "/Users/tangjian/Desktop/test/" + file.Name()
	// 	bytes, err := ioutil.ReadFile(path)
	// 	if err != nil {
	// 		println(err)
	// 		continue
	// 	}
	// 	str := string(bytes)
	// 	json.Unmarshal([]byte(s), &province)
	// 	println(str)
	// }
	now := time.Now().UTC()
	res, _ := time.ParseDuration("9999999s")
	now2 := now.Add(res)
	println(now2.String())
}
