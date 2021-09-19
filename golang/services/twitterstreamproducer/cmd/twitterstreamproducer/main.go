package main

import (
	"log"

	"github.com/clD11/go-whispers/golang/services/twitterstreamproducer/internal"
)

func main() {
	log.Println("starting twitter stream producer...")
	app := internal.App{}
	app.Initialize()
}
