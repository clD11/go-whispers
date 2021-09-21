package websocket_test

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"testing"
	"time"
)

func TestBroadcaster_Start(t *testing.T) {
	done := make(chan bool)
	go f(done)
	time.Sleep(time.Second * 60)
	done <- true
	log.Print("WAITING CONN CLOSE")
	time.Sleep(time.Second * 10)
}

func f(done chan bool) {
	u := url.URL{Scheme: "ws", Host: "localhost:8086", Path: "/subscribe"}
	conn, _, _ := websocket.DefaultDialer.Dial(u.String(), nil)
	defer conn.Close()

	for {
		select {
		case <-done:
			log.Print("closing conn TEST")
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			break
		default:
			_, msg, _ := conn.ReadMessage()
			log.Print("READ")
			log.Print(string(msg))
		}
	}

	log.Print("END")
}
