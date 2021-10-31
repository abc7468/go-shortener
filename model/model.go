package model

import (
	"fmt"

	db "github.com/abc7468/go-shortener/DB"
)

type ResponseData struct {
	OriginUrl string `json:"origin_url"`
}
type Shortener struct {
	ShortenerUrl string
	OriginalUrl  string
}

func GetAllUrls(port string) []Shortener {
	var val []Shortener
	urls := db.LoadAllUrls()
	for _, url := range urls {
		val = append(val, Shortener{fmt.Sprintf("localhost%s/%s", port, url[0]), url[1]})
	}
	return val
}
