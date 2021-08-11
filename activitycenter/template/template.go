package template

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
	// dir, err := ioutil.ReadDir("/Users/tangjian/Desktop/test")
	// if err != nil {
	// 	println(err)
	// 	return
	// }
	// for _, fs := range dir {

	// }
}
