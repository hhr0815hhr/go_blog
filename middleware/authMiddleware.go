package middleware

import (
	"blog/common"
	"blog/model"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")

		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			common.Response(ctx, 401, "权限不足", nil)
			ctx.Abort()
			return
		}
		tokenStr = tokenStr[7:]

		token, claims, err := common.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			common.Response(ctx, 401, "权限不足", nil)
			ctx.Abort()
			return
		}
		//验证通过
		userId := claims.UserId
		user := model.GetUserById(userId)
		if user.ID == 0 {
			common.Response(ctx, 401, "权限不足", nil)
			ctx.Abort()
			return
		}

		//用户存在
		ctx.Set("user", user)

		ctx.Next()
	}
}
