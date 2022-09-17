package entity

import (
	"database/sql"
	"fmt"
	"gameserver/mysql_pool"

	jsoniter "github.com/json-iterator/go"
)

type Entity interface {
	AddVersion()
	GetUserId() int
	Update()
	TableName() string
	GetTempId() (bool, int)
	Clone() Entity
}

func InsertUserData(userData *UserData) {
	Insert(userData.User, mysql_pool.MysqlPool)
	Insert(userData.GradeFrame, mysql_pool.MysqlPool)
	Insert(userData.Soul, mysql_pool.MysqlPool)
	for _, cter := range userData.Cters.CterMap {
		Insert(cter, mysql_pool.MysqlPool)
	}
}

func Insert(en Entity, sqlCon *sql.DB) error {
	_, tempId := en.GetTempId()
	jsonStr, jsonErr := jsoniter.Marshal(en)
	if jsonErr != nil {
		return jsonErr
	}
	var sqlStr string
	if tempId > 0 {
		sqlStr = fmt.Sprintf("insert into %s values(%d,%d,%s)", en.TableName(), en.GetUserId(), tempId, jsonStr)
	} else {
		sqlStr = fmt.Sprintf("insert into %s values(%d,%s)", en.TableName(), en.GetUserId(), jsonStr)
	}
	_, err := sqlCon.Exec(sqlStr)
	if err != nil {
		return err
	}
	return nil
}

func Update(en Entity, sqlCon *sql.DB) error {
	_, tempId := en.GetTempId()
	jsonStr, jsonErr := jsoniter.Marshal(en)
	if jsonErr != nil {
		return jsonErr
	}
	var sqlStr string
	if tempId > 0 {
		sqlStr = fmt.Sprintf("update %s set user_id=%d,temp_id=%d,content=%s where user_id=%d and temp_id=%d", en.TableName(), en.GetUserId(), tempId, jsonStr, en.GetUserId(), tempId)
	} else {
		sqlStr = fmt.Sprintf("update %s set user_id=%d,content=%s where user_id=%d", en.TableName(), en.GetUserId(), jsonStr, en.GetUserId())
	}
	_, err := sqlCon.Exec(sqlStr)
	if err != nil {
		return err
	}
	return nil
}

func QueryUserData(userId int) (*UserData, error) {
	var userInfo UserInfo
	var cters Characters
	cters.UserId = userId
	var gf GradeFrame
	var soul Soul

	userInfo.UserId = userId
	//查询UserInfo
	userInfoRes, err := Query(&userInfo, mysql_pool.MysqlPool)
	if err != nil {
		return nil, err
	}

	//查询cter
	var cter = Character{}
	cterRes, err := Query(&cter, mysql_pool.MysqlPool)
	if err != nil {
		return nil, err
	}
	for _, cterInterface := range cterRes {
		cter := cterInterface.(*Character)
		_, cterId := cter.GetTempId()
		cters.CterMap[cterId] = cter
	}

	gfRes, err := Query(&gf, mysql_pool.MysqlPool)
	if err != nil {
		return nil, err
	}

	soulRes, err := Query(&soul, mysql_pool.MysqlPool)
	if err != nil {
		return nil, err
	}

	userData := UserData{}

	userData.User = userInfoRes[0].(*UserInfo)
	userData.Cters = &cters
	userData.GradeFrame = gfRes[0].(*GradeFrame)
	userData.Soul = soulRes[0].(*Soul)
	return &userData, nil
}

func Query(en Entity, sqlCon *sql.DB) ([]interface{}, error) {
	var userId = en.GetUserId()
	var table = en.TableName()
	hasTempId, tempId := en.GetTempId()
	var sqlStr string
	var content string
	var err error
	sqlStr = fmt.Sprintf("select * from %s where user_id=%d", table, userId)

	rows, err := sqlCon.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	var errMess error
	resArray := []interface{}{}
	var i = 0
	for rows.Next() {
		if hasTempId {
			errMess = rows.Scan(&userId, &tempId, &content)
		} else {
			errMess = rows.Scan(&userId, &content)
		}
		if err != nil {
			return nil, errMess
		}
		//处理json
		bytes := []byte(content)
		res := en.Clone()
		err = jsoniter.Unmarshal(bytes, res)
		if err != nil {
			panic(err)
		}
		resArray[i] = res
		i += 1
	}
	return resArray, nil
}
