package main

import (
	"blog/controller/auth"
	"blog/controller/user"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	loginGroup := r.Group("/auth")
	{
		loginGroup.POST("/login", auth.Login)
		loginGroup.POST("/reg", auth.Reg)
		loginGroup.POST("/verifyName", auth.VerifyName)
		loginGroup.POST("/mail", auth.Mail)
	}
	userGroup := r.Group("/user")
	{
		userGroup.POST("info", middleware.AuthMiddleware(), user.GetUserInfo)
	}
}
