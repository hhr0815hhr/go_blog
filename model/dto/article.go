package dto

import (
	"blog/model"
	"strconv"
	"time"
)

type ArticleListDto struct {
	ID            string    `json:"article_id"`
	ArticleCate   string    `json:"article_category"`
	ArticleLabels []string  `json:"label"`
	Title         string    `json:"title"`
	ArticleImg    string    `json:"article_img"`
	Content       string    `json:"article_brief"`
	LikeNum       uint      `json:"like_Star"`
	CommentNum    uint      `json:"comment_num"`
	WatchNum      uint      `json:"visited"`
	CreatedAt     time.Time `json:"time"`
}

type RecentDto struct {
	ID        uint      `json:"article_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"time"`
}

type LabelDto struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func Labels(list []model.ArticleLabel) []LabelDto {
	var res = make([]LabelDto, len(list))
	for k, v := range list {
		res[k] = LabelDto{
			ID:   v.ID,
			Name: v.Name,
		}
	}
	return res
}

func ArticleList(list []model.Article) []ArticleListDto {
	var dto = make([]ArticleListDto, len(list))
	var content string
	for k, v := range list {
		content = v.Content
		if len(v.Content) >= 30 {
			content = v.Content[:30]
		}
		dto[k] = ArticleListDto{
			ID:            strconv.Itoa(int(v.ID)),
			ArticleCate:   model.GetCategoryNameById(v.ArticleCate),
			ArticleLabels: model.GetLabelInfoByIds(v.ArticleLabels),
			Title:         v.Title,
			ArticleImg:    v.ArticleImg,
			Content:       content,
			LikeNum:       v.LikeNum,
			CommentNum:    v.CommentNum,
			WatchNum:      v.WatchNum,
			CreatedAt:     v.CreatedAt,
		}
	}
	return dto
}

func RecentList(list []model.Article) []RecentDto {
	var res = make([]RecentDto, len(list))
	for k, v := range list {
		res[k] = RecentDto{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
		}
	}
	return res
}
