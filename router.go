package main

import (
	"blog/controller/auth"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/reg", auth.Reg)
		authGroup.POST("/verifyName", auth.VerifyName)
	}
}
