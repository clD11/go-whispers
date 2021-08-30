package main

import (
	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal"
	"log"
)

func main() {
	log.Println("Cryptocurrency Price Publisher")
	app := internal.NewApplication()
	app.Initialize()
}
