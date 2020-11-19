package main


type UserData struct {
	userId uint32
	cter Charcter
}

func newUserData(userId uint32) UserData{
	var cter = newCharcter(userId,1)
	var ud = UserData{userId: userId,cter: cter}
	return ud
}

type Charcter struct {
	UserId uint32
	CterId uint32
}

func  newCharcter(userId uint32,cterId uint32) Charcter{
	var cter = Charcter{UserId: 1,CterId: 1}
	return cter
}


