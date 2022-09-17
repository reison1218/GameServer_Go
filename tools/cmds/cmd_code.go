package cmds

/*游戏服务器需要处理的cmd*/
const (
	//最小范围
	GAME_MIN uint32 = 1000
	//心跳包
	HEARTBEAT uint32 = 1001
	//登陆
	LOGIN uint32 = 1002
	//卸载玩家数据
	UNLOADUSER uint32 = 9999
	//最大返回
	GAME_MAX uint32 = 10000
)
