package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/clD11/go-whispers/golang/services/shared"
)

type HealthResponse struct {
	Service   string `json:"service"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	log.Print("twitter stream producer health endpoint")
	shared.WriteResponse(w, http.StatusOK, HealthResponse{
		Service:   "twitter stream producer",
		Status:    "healthy",
		Timestamp: time.Now().String(),
	})
}
