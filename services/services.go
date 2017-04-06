package services

import (
	"github.com/astaxie/beego/orm"
)

type (
	DbService struct{
		Orm orm.Ormer
		DbName string
	}
)

const maxIdle = 30
const maxConn = 30

/**
connection pool
 */
var connectPull = make(map[string] *DbService)

func findServicesFromPull(dbname string)(dbInstance *DbService, res bool){
	dbInstance, ok := connectPull[dbname]
	if ok {
		return dbInstance, ok
	}

	return nil, ok
}

//get connect from here
func NewService(dbname string) *DbService {

	getInstance, ok := findServicesFromPull(dbname)

	if ok && getInstance != nil {
		return getInstance
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSource := "root:@/" + dbname + "?charset=utf8"

	//参数1        数据库的别名，用来在 ORM 中切换数据库使用
	//参数2        driverName
	//参数3        对应的链接字符串
	//参数4(可选)  设置最大空闲连接
	//参数5(可选)  设置最大数据库连接 (go >= 1.2)
	//user:password@/dbname

	orm.RegisterDataBase(dbname, "mysql", dataSource, maxIdle, maxConn)
	newOrm := orm.NewOrm()
	newInstance := &DbService{newOrm, dbname}

	//update connect pool
	connectPull[newInstance.DbName] = newInstance
	return newInstance
}

// Finish is called after the controller.
func (service *DbService) Finish() (err error) {
	return err
}