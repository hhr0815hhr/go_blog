package article

import (
	"blog/common"
	"blog/model"
	"blog/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageForm struct {
	Page uint `form:"page" json:"page" xml:"page" binding:"required"`
}

func GetArticleList(ctx *gin.Context) {
	var form PageForm
	if err := ctx.ShouldBind(&form); err != nil {
		common.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}
	page := form.Page
	if page <= 0 {
		common.Failed(ctx, -1, "获取文章列表失败")
		return
	}
	cnt := model.GetArticleCnt()
	if (page-1)*5 >= cnt {
		common.Failed(ctx, -1, "404,错误的页码数")
		return
	}
	list := model.GetArticleList(page)
	common.Success(ctx, "获取列表成功", gin.H{
		"count": cnt,
		"list":  dto.ArticleList(list),
	})
}

func GetRecentList(ctx *gin.Context) {
	list := model.GetRecentArticles()
	common.Success(ctx, "获取最近文章列表成功", gin.H{
		"list": dto.RecentList(list),
	})
	//common.Failed(ctx,-1,"失败")
}
