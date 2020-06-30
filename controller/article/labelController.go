package article

import (
	"blog/common"
	"blog/model"
	"blog/model/dto"
	"github.com/gin-gonic/gin"
)

func GetLabels(ctx *gin.Context) {
	labels := model.GetLabels()
	common.Success(ctx, "", gin.H{
		"list": dto.Labels(labels),
	})
}

func GetArticleLabelInfo(ctx *gin.Context) {
	common.Failed(ctx, -1, "获取数据失败")
}
