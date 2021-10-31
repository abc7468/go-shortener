package model

import db "github.com/abc7468/go-shortener/DB"

type Shortener struct {
	ShortenerUrl string
	OriginalUrl  string
}

func GetAllUrls() []Shortener {
	var val []Shortener
	urls := db.LoadAllUrls()
	for _, url := range urls {
		val = append(val, Shortener{url[0], url[1]})
	}
	return val
}
