package session

import "time"

type SessonInfo struct {
	GameId         int       "db:game_id"
	SessionId      int       "db:session_id"
	Round          int       "db:round"
	LastUpdateTime time.Time "db:last_update_time"
	NextUpdateTime time.Time "db:next_update_time"
}

type SessionMgr struct {
	SeasonMap map[uint32]SessonInfo
}

func (mgr SessionMgr) init() {

}
