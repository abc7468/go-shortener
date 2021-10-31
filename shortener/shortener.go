package shortener

import (
	"crypto/sha256"
	"fmt"

	db "github.com/abc7468/go-shortener/DB"
)

func DeleteUrlWithShortenUrl(shortenUrl string) error {
	err := db.DeleteUrl(shortenUrl)
	if err != nil {
		return err
	}
	return nil
}

func GetOriginWithShortenUrl(shortenUrl string) string {
	return db.LoadUrl(shortenUrl)
}

func SaveShortenUrl(originUrl string) {

	db.SaveURL(makeShortenUrl(originUrl), originUrl)
}

func makeShortenUrl(originUrl string) string {
	hash := goHash(originUrl)
	return hash[:8]
}

func goHash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}
