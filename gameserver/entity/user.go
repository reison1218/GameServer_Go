package entity

import "time"

type UserData struct {
	User       *UserInfo
	Soul       *Soul
	Cters      *Characters
	GradeFrame *GradeFrame
}

func (userData *UserData) UpdateLogin() {
	userData.User.Ol = true
	userData.User.LastLoginTime = time.Now().String()
	userData.User.AddVersion()
}
