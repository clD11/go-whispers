package handler

import (
	"net/http"
)

func GetLanding(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cryptocurrency Price Publisher - Using Kraken API"))
}
