package home

import (
	"github.com/abc7468/go-shortener/model"
	"github.com/gin-gonic/gin"
)

var port string

func ShortenerPage(c *gin.Context) {
	urls := model.GetAllUrls(port)

	c.HTML(200, "index.gohtml", gin.H{
		"title": "URL SHORTENER",
		"urls":  urls,
	},
	)
}

func Routing(portNum string, r *gin.Engine) {
	port = portNum
	r.LoadHTMLFiles("templates/index.gohtml")
	r.Static("css", "./templates/css")
	r.Static("js", "./templates/js")
	r.Static("imgs", "./templates/imgs")
	r.GET("/", ShortenerPage)
}
