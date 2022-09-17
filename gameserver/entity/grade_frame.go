package entity

/*头像框结构题*/
type GradeFrame struct {
	UserId      int   `json:"user_id"`      //玩家id
	GradeFrames []int `json:"grade_frames"` //头像框
	version     int   //版本号
}

func (gf *GradeFrame) AddVersion() {
	gf.version += 1
}
func (gf *GradeFrame) GetUserId() int {
	return gf.UserId
}
func (gf *GradeFrame) Update() {

}
func (gf *GradeFrame) TableName() string {
	return "t_u_grade_frame"
}
func (gf *GradeFrame) GetTempId() (bool, int) {
	return false, 0
}

func (ui *GradeFrame) Clone() Entity {
	return &GradeFrame{}
}
