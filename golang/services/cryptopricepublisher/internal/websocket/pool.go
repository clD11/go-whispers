package websocket

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// WIP
type Pool struct {
	conns map[uuid.UUID]*websocket.Conn
}

func New() *Pool {
	return &Pool{
		conns: make(map[uuid.UUID]*websocket.Conn),
	}
}

// upgrade and add connection to ws pool
func (p *Pool) Add(rw http.ResponseWriter, r *http.Request) error {
	u := websocket.Upgrader{}
	conn, err := u.Upgrade(rw, r, nil)
	if err != nil {
		return fmt.Errorf("wspool: could not upgrade conn %w", err)
	}
	p.conns[uuid.New()] = conn
	return nil
}
