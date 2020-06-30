package article

import (
	"blog/common"
	"blog/model"
	"github.com/gin-gonic/gin"
)

func GetCategories(ctx *gin.Context) {
	list := model.GetCategories()
	common.Success(ctx, "", gin.H{
		"list": list,
	})
}
