package main

import (
	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal"
	"log"
)

func main() {
	log.Println("starting cryptocurrency price publisher...")
	app := internal.App{}
	app.Initialize()
}
