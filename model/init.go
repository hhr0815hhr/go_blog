package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDb() {
	user := "root"
	pass := "123456"
	host := "192.168.0.115"
	port := "3306"
	sqlType := "mysql"
	dbName := "blog"
	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, dbName)
	var err error
	db, err = gorm.Open(sqlType, str)
	if err != nil {
		panic(fmt.Sprintf("数据库连接异常，err:%s", err.Error()))
	}
}

func GetDB() *gorm.DB {
	return db
}
