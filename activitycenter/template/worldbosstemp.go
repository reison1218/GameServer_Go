package template

type WorldBossTemplate struct {
	CterId   uint32
	KeepTime uint64
	MapIds   []uint32
}

type WorldBossTemplateMgr struct {
	templates map[uint32]WorldBossTemplate
}

func NewWorldBossTemplateMgr() WorldBossTemplateMgr {
	return WorldBossTemplateMgr{templates: make(map[uint32]WorldBossTemplate)}
}

func (mgr *WorldBossTemplateMgr) getById(id uint32) Template {
	return mgr.templates[id]
}

func (mgr *WorldBossTemplateMgr) init(temps []WorldBossTemplate) {
	for _, temp := range temps {
		mgr.templates[temp.CterId] = temp
	}
}
