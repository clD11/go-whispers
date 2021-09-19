package websockets

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// WIP
type WSPool struct {
	broadcast chan string
	conns     map[uuid.UUID]*websocket.Conn
}

func NewWsPool() *WSPool {
	return &WSPool{
		conns: make(map[uuid.UUID]*websocket.Conn),
	}
}

// creates ws connection and adds to pool
func (ws *WSPool) Connect(rw http.ResponseWriter, r *http.Request) error {
	u := websocket.Upgrader{}
	conn, err := u.Upgrade(rw, r, nil)
	if err != nil {
		return fmt.Errorf("wspool: could not upgrade conn %w", err)
	}
	ws.conns[uuid.New()] = conn
	return nil
}

// broadcast message to all connections
func (ws *WSPool) Broadcast(msg []byte) {
	for _, conn := range ws.conns {
		conn.WriteMessage(2, msg)
	}
}
