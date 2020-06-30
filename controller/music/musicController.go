package music

import (
	"blog/common"
	"github.com/gin-gonic/gin"
)

func GetMusic(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "1")
	common.Success(ctx, "", gin.H{
		"id":      id,
		"musicId": 25841337,
	})
}
