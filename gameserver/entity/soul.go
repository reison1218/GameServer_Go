package entity

type Soul struct {
	UserId  int   `json:"user_id"` //玩家id
	Souls   []int `json:"souls"`   //玩家拥有的灵魂
	version int   //版本号
}

func (soul *Soul) AddVersion() {
	soul.version += 1
}
func (soul *Soul) GetUserId() int {
	return soul.UserId
}
func (soul *Soul) Update() {

}
func (soul *Soul) TableName() string {
	return "t_u_soul"
}
func (soul *Soul) GetTempId() (bool, int) {
	return false, 0
}
func (ui *Soul) Clone() Entity {
	return &Soul{}
}
