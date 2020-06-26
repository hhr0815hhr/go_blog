package model

import (
	"blog/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDb() {
	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MysqlUser,
		config.MysqlPwd,
		config.MysqlHost,
		config.MysqlPort,
		config.MyDBName)
	var err error
	db, err = gorm.Open("mysql", str)
	if err != nil {
		panic(fmt.Sprintf("数据库连接异常，err:%s", err.Error()))
	}
	setDbPool()
	autoMigrate()
}

func autoMigrate() {
	db.SingularTable(true)
	db.AutoMigrate(&User{})
}

//连接池
func setDbPool() {
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
}

func GetDB() *gorm.DB {
	return db
}
