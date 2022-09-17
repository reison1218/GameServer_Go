package mysql_pool

import (
	"database/sql"
	"gameserver/config_helper"
	"gameserver/log"
	"tools/mysql_client"
)

var MysqlPool *sql.DB

func InitMysql() {
	addValue := config_helper.Configs.Configs["mysql"]
	mysqlStr := addValue.String()
	MysqlPool = mysql_client.InitMysql(mysqlStr)
	log.Logs.Info("mysql 初始化成功！")
}
