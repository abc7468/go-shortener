package rest

import (
	"fmt"
	"net/http"

	db "github.com/abc7468/go-shortener/DB"
	"github.com/abc7468/go-shortener/model"
	"github.com/abc7468/go-shortener/shortener"
	"github.com/gin-gonic/gin"
)

type url string

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost:8080/api/%s", u)
	return []byte(url), nil
}
func saveShortener(c *gin.Context) {
	data := model.ResponseData{}
	err := c.Bind(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	hash := shortener.MakeShorten(data.OriginUrl)
	db.SaveURL(hash, data.OriginUrl)
}
func documentation(c *gin.Context) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func Routing(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/", documentation)
	api.POST("/shortener", saveShortener)

}
