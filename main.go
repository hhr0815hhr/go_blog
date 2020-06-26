package main

import (
	"blog/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	SetRouter(r)
	model.InitDb()
	panic(r.Run())
}
