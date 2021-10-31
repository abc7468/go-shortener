package home

import (
	"github.com/abc7468/go-shortener/model"
	"github.com/gin-gonic/gin"
)

func ShortenerPage(c *gin.Context) {
	urls := model.GetAllUrls(":8080")

	c.HTML(200, "index.gohtml", gin.H{
		"title": "URL SHORTENER",
		"urls":  urls,
	},
	)
}

func Routing(r *gin.Engine) {
	r.LoadHTMLFiles("templates/index.gohtml")
	r.Static("css", "./templates/css")
	r.Static("js", "./templates/js")
	r.Static("imgs", "./templates/imgs")
	r.GET("/", ShortenerPage)
}
