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
	shortener.SaveShortenUrl(data.OriginUrl)
}

func findEndPoint(c *gin.Context) {
	data := c.Param("shortenUrl")
	originUrl := shortener.GetOriginWithShortenUrl(data)
	if originUrl == "" {
		c.String(http.StatusBadRequest, "%s는 잘못된 입력입니다.", data)
		return
	}
	c.Redirect(301, originUrl)
}

func deleteShortener(c *gin.Context) {
	data := model.ResponseData{}
	err := c.Bind(&data)
	if err != nil {
		fmt.Println(err)
	}
	db.DeleteUrl(data.OriginUrl[len(data.OriginUrl)-8:])
}
func documentation(c *gin.Context) {
	data := []urlDescription{
		{
			URL:         url("/api/shortener"),
			Method:      "POST",
			Description: "Make Shorten URL",
		},
		{
			URL:         url("/api/shortener"),
			Method:      "DELETE",
			Description: "DELETE Shorten URL",
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
	api.DELETE("/shortener", deleteShortener)
	r.GET("/:shortenUrl", findEndPoint)
}
