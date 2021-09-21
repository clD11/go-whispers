package websocket

import (
	"context"
	"fmt"
	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal/client"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Broadcaster struct {
	pool  *Pool
	close chan os.Signal
}

func NewBroadcaster() *Broadcaster {
	b := &Broadcaster{
		pool:  New(),
		close: make(chan os.Signal),
	}
	signal.Notify(b.close, os.Interrupt, syscall.SIGTERM)
	return b
}

func (b *Broadcaster) Subscribe(rw http.ResponseWriter, r *http.Request) error {
	err := b.pool.Connect(rw, r)
	if err != nil {
		return fmt.Errorf("broadcaster: could not add conn %w", err)
	}
	return nil
}

func (b *Broadcaster) Start() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context, pool *Pool) {
		k := client.NewKrakenClient()

		k.SubscribePublic(client.SubscribeMessage{
			Event: "subscribe",
			Pair:  []string{"XBT/USD"},
			Subscription: client.Subscription{
				Name: "ticker",
			},
		})

		for {
			select {
			case <-ctx.Done():
				log.Print("closing kraken client")
				k.Close()
				return
			default:
				k.ReadMessage()
			}
		}

	}(ctx, b.pool)

	<-b.close

	cancel()
}
