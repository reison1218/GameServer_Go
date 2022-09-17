package mysql_client

import (
	"database/sql"
)

func InitMysql(address string) *sql.DB {
	mysql, err := sql.Open("mysql", address)
	if err != nil {
		panic(err)
	}
	//最大连接数
	mysql.SetMaxOpenConns(128)
	//空闲连接数量
	mysql.SetMaxIdleConns(15)
	//最大连接周期，超过时间就close
	mysql.SetConnMaxLifetime(100 * 10)
	return mysql
}
