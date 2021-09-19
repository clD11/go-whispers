package websocket

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// WIP
type Pool struct {
	broadcast chan string
	conns     map[uuid.UUID]*websocket.Conn
}

func New() *Pool {
	return &Pool{
		conns: make(map[uuid.UUID]*websocket.Conn),
	}
}

// creates ws connection and adds to pool
func (p *Pool) Connect(rw http.ResponseWriter, r *http.Request) error {
	u := websocket.Upgrader{}
	conn, err := u.Upgrade(rw, r, nil)
	if err != nil {
		return fmt.Errorf("wspool: could not upgrade conn %w", err)
	}
	p.conns[uuid.New()] = conn
	return nil
}

// broadcast message to all connections
func (p *Pool) Broadcast(msg []byte) {
	for _, conn := range p.conns {
		conn.WriteMessage(2, msg)
	}
}
