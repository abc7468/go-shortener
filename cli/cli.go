package cli

import (
	"flag"
	"fmt"

	"github.com/abc7468/go-shortener/home"
	"github.com/abc7468/go-shortener/rest"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	port := flag.Int("port", 4000, "Set port of the server")
	flag.Parse()
	portNum := fmt.Sprintf(":%d", *port)

	fmt.Println(*port)
	home.Routing(portNum, router)
	rest.Routing(portNum, router)
	router.Run(portNum)

}
