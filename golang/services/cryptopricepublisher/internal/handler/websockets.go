package handler

import (
	"net/http"

	"github.com/clD11/go-whispers/golang/services/shared"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WSPool struct {
	broadcast chan string
	conns     map[uuid.UUID]*websocket.Conn
}

// creates ws connection and adds to pool
func (ws *WSPool) Connect(rw http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{}
	conn, err := u.Upgrade(rw, r, nil)
	if err != nil {
		shared.WriteErrorResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}
	ws.conns[uuid.New()] = conn
	rw.WriteHeader(http.StatusOK)
}
