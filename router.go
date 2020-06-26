package main

import (
	"blog/controller/auth"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	r.Group("/auth")
	{
		r.POST("/login", auth.Login)
	}
}
