package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

func InitDb() {
	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"), //config.MysqlUser,
		viper.GetString("mysql.pwd"),  //config.MysqlPwd,
		viper.GetString("mysql.host"), //config.MysqlHost,
		viper.GetString("mysql.port"), //config.MysqlPort,
		viper.GetString("mysql.db"),   //config.MyDBName)
	)
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
	db.AutoMigrate(&User{}, &Article{}, &ArticleCate{}, &ArticleLabel{})
}

//连接池
func setDbPool() {
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}

func GetDB() *gorm.DB {
	return db
}
