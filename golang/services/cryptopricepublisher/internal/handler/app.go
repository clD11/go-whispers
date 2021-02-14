package handler

import (
	"log"
	"net/http"
)

func GetLanding(w http.ResponseWriter, r *http.Request) {
	log.Println("Splash!")
	w.WriteHeader(200)
	w.Write([]byte("Cryptocurrency Price Publisher - Using Kraken API"))
}
