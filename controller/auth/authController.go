package auth

import (
	"blog/common"
	"blog/model"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Account string `form:"user" json:"user" xml:"user" binding:"required"`
	Pwd     string `form:"password" json:"password" xml:"password" binding:"required"`
	Email   string `form:"email" json:"email" xml:"email"`
	Code    string `form:"code" json:"code" xml:"code"`
}

type RegForm struct {
	Account string `form:"user" json:"user" xml:"user" binding:"required"`
	Pwd     string `form:"password" json:"password" xml:"password" binding:"required"`
	Email   string `form:"email" json:"email" xml:"email"`
	Code    string `form:"code" json:"code" xml:"code"`
}

type VerifyForm struct {
	Account string `form:"user" json:"user" xml:"user" binding:"required"`
}

type MailForm struct {
	Email string `form:"email" json:"email" xml:"email" binding:"required"`
}

func Login(ctx *gin.Context) {
	var form LoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		common.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if form.Code != "" {
		ctx.Redirect(http.StatusPermanentRedirect, "/auth/reg")
		return
	}

	user := model.GetUserByAccount(form.Account)

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
	name := form.Account //ctx.PostForm("name")
	if !verifyParam("name", name) {
		common.Response(ctx, 422, "用户名长度8~20!", nil)
		return
	}
	isExist := model.IsExist(name)
	if isExist {
		common.Response(ctx, 201, "已经存在的用户名", nil)
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
	name := form.Account //ctx.PostForm("name")
	pwd := form.Pwd      //ctx.PostForm("pwd")
	email := form.Email  //ctx.PostForm("phone")
	if !verifyParam("name", name) {
		common.Response(ctx, 422, "用户名长度8~20!", nil)
		return
	}
	if !verifyParam("pwd", pwd) {
		common.Response(ctx, 422, "密码长度8~20!", nil)
	}
	if !verifyParam("email", email) {
		common.Response(ctx, 422, "请输入正确的邮箱", nil)
		return
	}
	isExist := model.IsExist(name)
	if isExist {
		common.Response(ctx, 201, "已经存在的用户名", nil)
		return
	}
	pwd, _ = util.BCrypt(pwd)
	err, user := model.RegUser(name, pwd, email)
	if err != nil {
		common.Response(ctx, 422, "注册失败,err: "+err.Error(), nil)
		return
	}
	token, _ := common.ReleaseToken(user)
	common.Success(ctx, "注册成功", gin.H{
		"token": token,
	})
	return
}

func Mail(ctx *gin.Context) {
	var form MailForm
	if err := ctx.ShouldBind(&form); err != nil {
		common.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err := common.Mail(form.Email, "ahaha", "1234", "html")
	var msg = "发送成功"
	if err != nil {
		msg = err.Error()
		common.Response(ctx, 111, err.Error(), nil)
	}
	common.Success(ctx, msg, nil)
}

func verifyParam(key string, value interface{}) (b bool) {
	switch key {
	case "name", "pwd":
		if len(value.(string)) <= 20 && len(value.(string)) >= 8 {
			b = true
		}
	case "email":
		if util.RegexpEmail(value.(string)) != nil {
			b = true
		}
	default:
	}
	return
}
