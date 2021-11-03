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

var port string

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}
func saveShortener(c *gin.Context) {
	data := model.ResponseData{}
	err := c.Bind(&data)
	if err != nil {
		fmt.Println(err)
	}
	errNum := shortener.SaveShortenUrl(data.OriginUrl)
	switch errNum {
	case shortener.ErrNotRightUrl:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "옳지 않은 입력값",
		})
	case shortener.ErrNotUniqueUrl:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "중복된 입력값",
		})
	case shortener.ErrNotUniqueHash:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "해시 충돌. 다른 URL을 입력하세요.",
		})
	}
}

func findEndPoint(c *gin.Context) {
	data := c.Param("shortenUrl")
	originUrl := shortener.GetOriginWithShortenUrl(data)
	if data != "favicon.ico" && originUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%s는 저장된 값이 아닙니다.", data),
		})

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
	c.Header("Content-Type", "application/json")
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "Shortener Main Page",
		},
		{
			URL:         url("/api"),
			Method:      "GET",
			Description: "Document Page",
		},
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
		{
			URL:         url("/{Hash}"),
			Method:      "GET",
			Description: "Go Original URL with Shorten URL if Hash stored in DB",
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func Routing(portNum string, r *gin.Engine) {
	port = portNum
	api := r.Group("/api")
	api.GET("/", documentation)
	api.POST("/shortener", saveShortener)
	api.DELETE("/shortener", deleteShortener)
	r.GET("/:shortenUrl", findEndPoint)
}
