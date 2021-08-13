package template

type WorldBossTemplate struct {
	CterId   int   `json:"cter_id"`
	KeepTime int64 `json:"keep_time"`
	MapIds   []int `json:"map_ids"`
	RobotId  int   `json:"robot_id"`
}

type WorldBossTemplateMgr struct {
	templates map[int]WorldBossTemplate
}

func (mgr *WorldBossTemplateMgr) GetFirst() WorldBossTemplate {
	return mgr.templates[2001]
}

func (mgr *WorldBossTemplateMgr) GetNext(cterId int) WorldBossTemplate {
	res := mgr.templates[2001]
	for _, i := range mgr.templates {
		if i.CterId != cterId {
			return i
		}
	}
	return res
}

func NewWorldBossTemplateMgr() WorldBossTemplateMgr {
	return WorldBossTemplateMgr{templates: make(map[int]WorldBossTemplate)}
}

func (mgr *WorldBossTemplateMgr) GetById(id int) WorldBossTemplate {
	return mgr.templates[id]
}

func (mgr *WorldBossTemplateMgr) IsEmpty() bool {
	return len(mgr.templates) == 0
}

func (mgr *WorldBossTemplateMgr) Clear() {
	mgr.templates = make(map[int]WorldBossTemplate)
}

func (mgr *WorldBossTemplateMgr) init(temps []WorldBossTemplate) {
	for _, temp := range temps {
		mgr.templates[temp.CterId] = temp
	}
}
