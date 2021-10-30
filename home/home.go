package home

import (
	"github.com/abc7468/go-shortener/model"
	"github.com/gin-gonic/gin"
)

func ShortenerPage(c *gin.Context) {
	urls := model.GetAllUrls()

	c.HTML(200, "index.gohtml", gin.H{
		"title": "URL SHORTENER",
		"urls":  urls,
	},
	)
}

func Start() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.gohtml")
	r.Static("css", "./templates/css")
	r.Static("js", "./templates/js")
	r.Static("imgs", "./templates/imgs")
	r.GET("/", ShortenerPage)
	r.Run()
}
