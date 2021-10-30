package main

import (
	db "github.com/abc7468/go-shortener/DB"
	"github.com/abc7468/go-shortener/home"
)

func main() {
	db.InitDB()
	home.Start()
}
