package main

import (
	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal"
	"log"
)

func main() {
	log.Println("starting crypto price publisher")
	app := internal.NewApplication()
	app.Start(":8080")
}
