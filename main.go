package main

import (
	_ "blog/config"
	"blog/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//r := gin.Default()
	r := gin.New()
	r.Use(AccessControlAllowOrigin())
	SetRouter(r)
	model.InitDb()
	panic(r.Run(":" + viper.GetString("server.port")))
}
