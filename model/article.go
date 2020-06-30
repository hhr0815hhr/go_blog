package model

import (
	"strings"
	"time"
)

type Article struct {
	ID            uint   `gorm:"primary_key"`
	ArticleCate   uint   `gorm:"index:idx_cate"`
	ArticleLabels string `gorm:"size:20;not null"`
	Title         string `gorm:"size:20;not null"`
	ArticleImg    string `gorm:"size:50"`
	Content       string `gorm:"type=text;not null"`
	LikeNum       uint
	CommentNum    uint
	WatchNum      uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ArticleCate struct {
	ID       uint   `gorm:"primary_key"`
	PId      uint   `gorm:"default:0;not null"`
	CateName string `gorm:"size:20;not null"`
	Count    uint   `gorm:"default:0;not null"`
}

type ArticleLabel struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:20;not null"`
}

var (
	article    Article
	articles   []Article
	labels     []ArticleLabel
	categories []ArticleCate
	category   ArticleCate
)

func GetArticleCnt() uint {
	var count uint
	db.Table("article").Count(&count)
	return count
}

func GetArticleList(page uint) []Article {
	db.Select("id,article_cate,article_labels,title,article_img,like_num,watch_num,comment_num,created_at,content").Order("created_at desc").Limit(5).Offset(page - 1).Find(&articles)
	return articles
}

func GetRecentArticles() []Article {
	db.Select("id,title,created_at").Order("created_at desc").Limit(2).Find(&articles)
	return articles
}

func GetLabels() []ArticleLabel {
	db.Find(&labels)
	return labels
}

func GetCategories() []ArticleCate {
	db.Find(&categories)
	return categories
}

func GetCategoryNameById(cateId uint) string {
	db.Where("id=?", cateId).Find(&category)
	return category.CateName
}

func GetLabelInfoByIds(labelIds string) []string {
	ids := strings.Split(labelIds, ",")
	db.Where("id in (?)", ids).Find(&labels)
	var res = make([]string, len(labels))
	for k, v := range labels {
		res[k] = v.Name
	}
	return res
}
