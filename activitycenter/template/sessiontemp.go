package template

type SessionTemplate struct {
	Id      uint32
	Element uint32
}

func (temp *SessionTemplate) getId() uint32 { return temp.Id }

type SessionTemplateMgr struct {
	templates map[uint32]SessionTemplate
}

func NewSessionTemplateMgr() SessionTemplateMgr {
	return SessionTemplateMgr{templates: make(map[uint32]SessionTemplate)}
}

func (mgr *SessionTemplateMgr) getById(id uint32) Template {
	return mgr.templates[id]
}

func (mgr *SessionTemplateMgr) isEmpty() bool {
	return len(mgr.templates) == 0
}

func (mgr *SessionTemplateMgr) clear() {
	mgr.templates = make(map[uint32]SessionTemplate)
}

func (mgr *SessionTemplateMgr) init(temps []SessionTemplate) {
	for _, temp := range temps {
		mgr.templates[temp.Id] = temp
	}
}
