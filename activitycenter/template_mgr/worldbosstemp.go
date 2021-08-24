package template_mgr

import "container/list"

type WorldBossTemplate struct {
	CterId   int   `json:"cter_id"`
	KeepTime int64 `json:"keep_time"`
	MapIds   []int `json:"map_ids"`
	RobotId  int   `json:"robot_id"`
}

type WorldBossTemplateMgr struct {
	templates *list.List
}

func (mgr *WorldBossTemplateMgr) GetFirst() *WorldBossTemplate {
	res := mgr.templates.Front().Value.(WorldBossTemplate)
	return &res
}

func (mgr *WorldBossTemplateMgr) GetNext(cterId int) *WorldBossTemplate {
	for i := mgr.templates.Front(); i != nil; i = i.Next() {
		temp := i.Value.(WorldBossTemplate)
		if temp.CterId != cterId {
			continue
		}
		mgr.templates.MoveToBack(i)
		return &temp
	}
	return nil
}

func NewWorldBossTemplateMgr() WorldBossTemplateMgr {
	return WorldBossTemplateMgr{templates: list.New()}
}

func (mgr *WorldBossTemplateMgr) GetById(id int) *WorldBossTemplate {
	for i := mgr.templates.Front(); i != nil; i = i.Next() {
		temp := i.Value.(WorldBossTemplate)
		if temp.CterId != id {
			continue
		}
		return &temp
	}
	return nil
}

func (mgr *WorldBossTemplateMgr) IsEmpty() bool {
	return mgr.templates.Len() == 0
}

func (mgr *WorldBossTemplateMgr) Clear() {
	mgr.templates = list.New()
}

func (mgr *WorldBossTemplateMgr) init(temps []WorldBossTemplate) {
	for _, temp := range temps {
		mgr.templates.PushFront(temp)
	}
}
