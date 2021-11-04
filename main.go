package main

import (
	db "github.com/abc7468/go-shortener/DB"
	"github.com/abc7468/go-shortener/cli"
)

func main() {
	db.InitDB()
	defer db.Close()
	cli.Start()
}
