package entity

type Characters struct {
	UserId  int                //玩家id
	CterMap map[int]*Character //角色数据
	Version int                //版本号
}
type Character struct {
	UserId        int     `json:"user_id"`         //玩家id
	TempId        int     `json:"temp_id"`         //角色id
	UseTimes      int     `json:"use_times"`       //使用次数
	Skills        []Group `json:"skills"`          //技能组
	LastUseSkills []int   `json:"last_use_skills"` //上次使用的技能
	version       int     //版本号
}

func (cter *Character) AddVersion() {
	cter.version += 1
}

type Group struct {
	Skills []int `json:"skills"` //技能数组
}

func (ui *Character) GetUserId() int {
	return ui.UserId
}

func (ui *Character) Update() {
	//更新到数据库
	if ui.version > 0 {

	}
	//更新完了就清空版本号
	ui.version = 0
}

func (ui *Character) TableName() string {
	return "t_u_character"
}
func (ui *Character) GetTempId() (bool, int) {
	return true, ui.TempId
}
func (ui *Character) Clone() Entity {
	return &Character{}
}

func QueryCharacter() Character {
	return Character{}
}
