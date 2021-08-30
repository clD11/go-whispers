package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetLanding(w http.ResponseWriter, r *http.Request) {
	log.Println("Splash!")

	res, err := json.Marshal(Response{
		Data: "cryptocurrency price publisher v1",
	})

	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("bad"))
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

type Response struct {
	Data string `json:"data"`
}
