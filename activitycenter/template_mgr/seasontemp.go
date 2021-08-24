package template_mgr

import "container/list"

type SeasonTemplate struct {
	Id       int   `json:"id"`
	Element  int   `json:"element"`
	KeepTime int64 `json:"keep_time"`
}

type SeasonTemplateMgr struct {
	templates *list.List
}

func (mgr *SeasonTemplateMgr) GetNext(id int) *SeasonTemplate {
	for i := mgr.templates.Front(); i != nil; i = i.Next() {
		temp := i.Value.(SeasonTemplate)
		if temp.Id != id {
			continue
		}
		mgr.templates.MoveToBack(i)
		return &temp
	}
	return nil
}

func NewSeasonTemplateMgr() SeasonTemplateMgr {
	return SeasonTemplateMgr{templates: list.New()}
}

func (mgr *SeasonTemplateMgr) GetById(id int) *SeasonTemplate {
	for i := mgr.templates.Front(); i != nil; i = i.Next() {
		temp := i.Value.(SeasonTemplate)
		if temp.Id != id {
			continue
		}
		return &temp
	}
	return nil
}

func (mgr *SeasonTemplateMgr) IsEmpty() bool {
	return mgr.templates.Len() == 0
}

func (mgr *SeasonTemplateMgr) Clear() {
	mgr.templates = list.New()
}

func (mgr *SeasonTemplateMgr) init(temps []SeasonTemplate) {
	for _, temp := range temps {
		mgr.templates.PushFront(temp)
	}
}
