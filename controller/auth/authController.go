package auth

import (
	"blog/common"
	"blog/model"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Phone int64  `form:"phone" json:"phone" xml:"phone" binding:"required"`
	Pwd   string `form:"pwd" json:"pwd" xml:"pwd" binding:"required"`
}

type RegForm struct {
	Phone int64  `form:"phone" json:"phone" xml:"phone" binding:"required"`
	Pwd   string `form:"pwd" json:"pwd" xml:"pwd" binding:"required"`
	Name  string `form:"name" json:"name" xml:"name" binding:"required"`
}

type VerifyForm struct {
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
}

func Login(ctx *gin.Context) {
	var form LoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		common.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//pwd := ctx.PostForm("pwd")
	//phone := ctx.PostForm("phone")
	//if !verifyParam("phone", phone) {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code": 422,
	//		"msg":  "请输入正确的手机号",
	//	})
	//	return
	//}
	//if !verifyParam("pwd", pwd) {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code": 422,
	//		"msg":  "密码长度8~20!",
	//	})
	//	return
	//}
	user := model.GetUserByPhone(form.Phone)

	if user.ID == 0 || !util.BCryptVerify([]byte(form.Pwd), []byte(user.Pass)) {
		common.Failed(ctx, http.StatusUnauthorized, "账号或密码错误")
		return
	}
	token, _ := common.ReleaseToken(user)
	common.Success(ctx, "登陆成功", gin.H{"token": token})
	return
}

func VerifyName(ctx *gin.Context) {
	var form VerifyForm
	if err := ctx.ShouldBind(&form); err != nil {
		common.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}
	name := form.Name //ctx.PostForm("name")
	if !verifyParam("name", name) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, "用户名长度8~20!", nil)
		return
	}
	isExist := model.IsExist(name)
	if isExist {
		common.Response(ctx, http.StatusCreated, 201, "已经存在的用户名", nil)
		return
	}
	common.Success(ctx, "可以使用的用户名", nil)
}

func Reg(ctx *gin.Context) {
	var form RegForm
	if err := ctx.ShouldBind(&form); err != nil {
		common.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}
	name := form.Name   //ctx.PostForm("name")
	pwd := form.Pwd     //ctx.PostForm("pwd")
	phone := form.Phone //ctx.PostForm("phone")
	if !verifyParam("name", name) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, "用户名长度4~20!", nil)
		return
	}
	if !verifyParam("pwd", pwd) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, "密码长度8~20!", nil)
	}
	if !verifyParam("phone", phone) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, "请输入正确的手机号", nil)
		return
	}
	pwd, _ = util.BCrypt(pwd)
	err := model.RegUser(name, pwd, phone)
	if err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, "注册失败,err: "+err.Error(), nil)
		return
	}
	common.Success(ctx, "注册成功", nil)
	return
}

func verifyParam(key string, value interface{}) (b bool) {
	switch key {
	case "name", "pwd":
		if len(value.(string)) <= 20 && len(value.(string)) >= 8 {
			b = true
		}
	case "phone":
		if len(util.Int2String(int(value.(int64)))) == 11 {
			b = true
		}
	default:
	}
	return
}
