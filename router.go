package main

import (
	"blog/controller/article"
	"blog/controller/auth"
	"blog/controller/music"
	"blog/controller/user"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	testGroup := r.Group("/test")
	{
		testGroup.GET("/test", auth.Login)
	}
	loginGroup := r.Group("/auth")
	{
		loginGroup.POST("/login", auth.Login)
		loginGroup.POST("/reg", auth.Reg)
		loginGroup.POST("/verifyName", auth.VerifyName)
		loginGroup.POST("/mail", auth.Mail)
	}
	userGroup := r.Group("/user")
	{
		userGroup.GET("info", middleware.AuthMiddleware(), user.GetUserInfo)
	}
	articleGroup := r.Group("/article")
	{
		articleGroup.GET("/getList", article.GetArticleList)
		articleGroup.GET("getRecentList", article.GetRecentList)

		articleGroup.GET("/getLabels", article.GetLabels)
		articleGroup.POST("/getLabelInfo", article.GetArticleLabelInfo)

		articleGroup.GET("getCategories", article.GetCategories)
	}

	musicGroup := r.Group("/music")
	{
		musicGroup.GET("getMusic/:id", music.GetMusic)
	}
}
