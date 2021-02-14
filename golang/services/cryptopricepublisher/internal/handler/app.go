package handler

import (
	"log"
	"net/http"
)

func GetLanding(w http.ResponseWriter, r *http.Request) {
	log.Println("Splash!")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"data\":\"cryptocurrency price publisher v1\"}"))
}
