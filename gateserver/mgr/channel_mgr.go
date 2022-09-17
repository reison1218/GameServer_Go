package mgr

import (
	"net"
)

var CM ChannelMgr

type SH interface {
	Write(bytes []byte)
}

type ChannelMgr struct {
	//客户端tcp链接
	GameClientChannel *net.TCPConn
	//tcp链接
	UserChannel map[uint32]*GateUser
	//token,user_id
	Channel map[uint32]uint32
	//临时链接
	TempChannel map[uint32]SH
}

/*有效会话玩家结构体*/
type GateUser struct {
	SS SH
}
