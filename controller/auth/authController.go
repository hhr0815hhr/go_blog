package auth

import (
	"blog/model"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	pwd := ctx.PostForm("pwd")
	phone := ctx.PostForm("phone")
	if !verifyParam("phone", phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "请输入正确的手机号",
		})
		return
	}
	if !verifyParam("pwd", pwd) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码长度8~20!",
		})
	}
	user := model.GetUserByPhone(util.String2Int64(phone))

	if user.ID == 0 || !util.BCryptVerify([]byte(pwd), []byte(user.Pass)) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "账号或密码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登陆成功",
	})
	return
}

func VerifyName(ctx *gin.Context) {
	name := ctx.PostForm("name")
	if !verifyParam("name", name) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名长度8~20!",
		})
		return
	}
	isExist := model.IsExist(name)
	if isExist {
		ctx.JSON(http.StatusCreated, gin.H{
			"code": 201,
			"msg":  "已经存在的用户名",
		})
		return
	}
}

func Reg(ctx *gin.Context) {
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")
	phone := ctx.PostForm("phone")
	if !verifyParam("name", name) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名长度8~20!",
		})
		return
	}
	if !verifyParam("pwd", pwd) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码长度8~20!",
		})
	}
	if !verifyParam("phone", phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "请输入正确的手机号",
		})
		return
	}
	if !verifyParam("name", name) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名长度8~20!",
		})
		return
	}
	err := model.RegUser(name, pwd, util.String2Int64(phone))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "注册失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
	return
}

func verifyParam(key string, value interface{}) (b bool) {
	switch key {
	case "name", "pwd":
		if len(value.(string)) <= 20 && len(value.(string)) >= 8 {
			b = true
		}
	case "phone":
		if len(value.(string)) == 11 {
			b = true
		}
	default:
	}
	return
}
