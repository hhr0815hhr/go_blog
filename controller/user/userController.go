package user

import (
	"blog/common"
	"blog/model"
	"blog/model/dto"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	common.Success(ctx, "获取数据成功", gin.H{"user": dto.User(user.(*model.User))})
	return
}
