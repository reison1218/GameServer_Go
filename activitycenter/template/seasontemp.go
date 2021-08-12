package template

type SeasonTemplate struct {
	Id      int `json:"id"`
	Element int `json:"element"`
}

type SeasonTemplateMgr struct {
	templates map[int]SeasonTemplate
}

func NewSeasonTemplateMgr() SeasonTemplateMgr {
	return SeasonTemplateMgr{templates: make(map[int]SeasonTemplate)}
}

func (mgr *SeasonTemplateMgr) GetById(id int) SeasonTemplate {
	return mgr.templates[id]
}

func (mgr *SeasonTemplateMgr) IsEmpty() bool {
	return len(mgr.templates) == 0
}

func (mgr *SeasonTemplateMgr) Clear() {
	mgr.templates = make(map[int]SeasonTemplate)
}

func (mgr *SeasonTemplateMgr) init(temps []SeasonTemplate) {
	for _, temp := range temps {
		mgr.templates[temp.Id] = temp
	}
}
