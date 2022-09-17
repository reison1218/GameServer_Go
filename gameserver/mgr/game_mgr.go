package mgr

import (
	"fmt"
	"gameserver/entity"
	"gameserver/log"
	"tools/utils"
)

type SH interface {
	Write(bytes []byte)
}

var GM *GameMgr = &GameMgr{}

type GameMgr struct {
	Users   map[int]*entity.UserData
	CmdMap  map[int]func(gm *GameMgr, packet *utils.Packet)
	TcpConn SH
}

func (gm *GameMgr) LoginOff(userId int) {
	delete(gm.Users, userId)
}

func (gm *GameMgr) invok(packet *utils.Packet) {
	cmd := packet.Des.Cmd
	fn := gm.CmdMap[int(cmd)]
	if fn == nil {
		str := fmt.Sprintf("can not find func for cmd:%d", cmd)
		log.Logs.Error(str)
		return
	}
	fn(gm, packet)
}
