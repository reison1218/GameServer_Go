package entity

import (
	"database/sql"
	"fmt"
)

/*玩家基础数据*/
type UserInfo struct {
	UserId          int    `json:"user_id"`           //玩家id
	Ol              bool   `json:"ol"`                //当前是否在线
	NickName        string `json:"nick_name"`         //昵称
	Grade           int    `json:"grade"`             //玩家等级
	Soul            int    `json:"soul"`              //玩家灵魂
	GradeFrame      int    `json:"grade_frame"`       //grade像框
	LastLoginTime   string `json:"last_login_time"`   //最近登陆时间
	LastOffTime     string `json:"last_off_time"`     //最近离线时间
	LastCharacter   int    `json:"last_character"`    //最近使用的角色
	TotalOnlineTime int64  `json:"total_online_time"` //总在线时间
	version         int    //版本号（用于更新数据库）
}

func (ui *UserInfo) GetUserId() int {
	return ui.UserId
}

func (ui *UserInfo) Update() {
	//更新到数据库
	if ui.version > 0 {

	}
	//更新完了就清空版本号
	ui.version = 0
}

func (ui *UserInfo) TableName() string {
	return "t_u_player"
}
func (ui *UserInfo) GetTempId() (bool, int) {
	return false, 0
}
func (ui *UserInfo) AddVersion() {
	ui.version = +1
}

func (ui *UserInfo) Clone() Entity {
	return &UserInfo{}
}

func QueryUserInfo(userId int, sqlcon *sql.DB) (UserInfo, error) {
	sqlStr := fmt.Sprintf("select * from t_u_player where user_id=%d", userId)
	row := sqlcon.QueryRow(sqlStr)
	userInfo := UserInfo{}
	content := ""
	err := row.Scan(&userInfo.UserId, &content)
	if err != nil {
		return userInfo, err
	}

	return UserInfo{}, nil
}
