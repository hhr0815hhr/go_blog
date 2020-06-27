package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//定义统一返回格式

func Response(ctx *gin.Context, code int, msg string, data gin.H) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Success(ctx *gin.Context, msg string, data gin.H) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": SuccessCode,
		"msg":  msg,
		"data": data,
	})
}

func Failed(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
