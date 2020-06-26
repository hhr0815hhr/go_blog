package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登陆成功",
	})
	return
}
