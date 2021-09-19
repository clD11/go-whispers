package handler

import (
	"log"
	"net/http"

	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal/websocket"
	"github.com/clD11/go-whispers/golang/services/shared"
)

type subscribePriceUpdateResponse struct {
	success string `json:"success"`
}

// GET subscribe client for price updates
func SubscribePriceUpdate(wsPool *websocket.Pool, rw http.ResponseWriter, r *http.Request) {
	log.Print("adding connection to pool")
	err := wsPool.Connect(rw, r)
	if err != nil {
		shared.WriteErrorResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}
	shared.WriteWrappedResponse(rw, http.StatusOK, subscribePriceUpdateResponse{
		success: "connected",
	})
}
