package shortener

import (
	"crypto/sha256"
	"fmt"
	"regexp"

	db "github.com/abc7468/go-shortener/DB"
)

type errNum int

const (
	OK errNum = iota
	ErrNotRightUrl
	ErrNotUniqueUrl
	ErrNotUniqueHash
)

func SaveShortenUrl(originUrl string) errNum {
	shortenUrl, err := isCanSaveOriginUrl(originUrl)
	if err != OK {
		return err
	}
	db.SaveURL(shortenUrl, originUrl)
	return OK
}

//입력 URL이 저장 가능한 데이터인지 확인
func isCanSaveOriginUrl(originUrl string) (string, errNum) {
	if !isRightURL(originUrl) {
		return "", ErrNotRightUrl
	}
	shortenUrl := makeShortenUrl(originUrl)
	//해당 해시값이 존재한다면 저장하지 못하는 값
	if check := db.LoadUrl(shortenUrl); check != "" {
		if check == originUrl { // 이미 저장된 originUrl
			return "", ErrNotUniqueUrl
		}
		// Hash값이 겹치는 데이터
		return "", ErrNotUniqueHash
	}
	return shortenUrl, OK
}

//입력 URL이 옳은 형태의 URL인지 정규식을 사용해 check
func isRightURL(originUrl string) bool {
	r, _ := regexp.Compile("^(file|gopher|news|nntp|telnet|https?|ftps?|sftp)://([a-z0-9-]+\\.)+[a-z0-9]{2,4}.*$")
	return r.MatchString(originUrl)
}

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

func makeShortenUrl(originUrl string) string {
	hash := goHash(originUrl)
	return hash[:8]
}

func goHash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}
