package app

import (
	db "github.com/abc7468/go-shortener/DB"
	"github.com/abc7468/go-shortener/home"
	"github.com/abc7468/go-shortener/rest"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Start() {
	db.InitDB()
	defer db.Close()
	router = gin.Default()
	home.Routing(router)
	rest.Routing(router)
	router.Run(":8080")
}
