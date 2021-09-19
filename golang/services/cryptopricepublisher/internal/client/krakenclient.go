package client

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

const (
	krakenWSUrl = "ws.kraken.com"
	scheme      = "wss"
)

type KrakenClient struct {
	conn *websocket.Conn
}

type SubscribeMessage struct {
	Event        string       `json:"event"`
	ReqID        int          `json:"reqid,omitempty"`
	Pair         []string     `json:"pair"`
	Subscription Subscription `json:"subscription"`
}

type Subscription struct {
	Depth    int    `json:"depth,omitempty"`
	Interval int    `json:"interval,omitempty"`
	Name     string `json:"name"`
	Token    string `json:"token,omitempty"`
}

type TickerMessage struct {
	ask                        []string `json:"a"`
	bid                        []string `json:"b"`
	lastTradeClosed            []string `json:"c"`
	volume                     []string `json:"v"`
	volumeWeightedAveragePrice []string `json:"p"`
	numberOfTrades             []string `json:"t"`
	low                        []string `json:"l"`
	high                       []string `json:"h"`
	todayOpenPrice             string   `json:"o"`
}

// return a new kraken client and opens a new connection
func newKrakenClient() *KrakenClient {
	url := url.URL{Scheme: scheme, Host: krakenWSUrl}
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		fmt.Errorf("krakenclient: could not create new client %w", err)
	}

	return &KrakenClient{
		conn: conn,
	}
}

var count = 0

// subscribes to a public channel by sending a subscribe message
func (k *KrakenClient) SubscribePublic(subscribeMessage SubscribeMessage) error {
	msg, err := json.Marshal(subscribeMessage)
	if err != nil {
		return fmt.Errorf("krakenclient: could not serialize msg %w", err)
	}

	err = k.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return fmt.Errorf("krakenclient: could not subscribe to channel %w", err)
	}

	input := make(chan string)

	go func(input chan string) {
		for {
			_, _, err := k.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
		}
	}(input)

	return nil
}

func (k *KrakenClient) Close() {
	k.conn.Close()
}
