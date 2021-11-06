package util

import (
	"PushServer/conf"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Db为mysql连接句柄
var (
	Db  *sql.DB
	err error
)

//再调用此包时，自动调用init函数
func init() {
	sqlInit := conf.Conf.MysqlYaml.User + ":" + conf.Conf.MysqlYaml.Passowrd + "@tcp(" + conf.Conf.MysqlYaml.Addr + ")/" + conf.Conf.MysqlYaml.DbName
	// fmt.Println(sqlInit)
	Db, err = sql.Open("mysql", sqlInit)
	if err != nil {
		fmt.Println("连接数据库失败，请检查数据库是否开启，数据库是否选择正确!")
		panic(err.Error())
	}
	//查询是否创建表单，如果没用创建表单就执行sql文件创建表单
}
