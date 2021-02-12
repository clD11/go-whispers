package internal

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

const (
	krakenWSUrl = "ws.kraken.com"
)

type KrakenClient struct {
	conn *websocket.Conn
}

type KrakenResponseWrapper struct {
	result string   `json:"result"`
	error  []string `json:"error"`
}

type SubscribeEvent struct {
	Event        string       `json:"event"`
	Reqid        int          `json:"reqid,omitempty"`
	Pair         []string     `json:"pair"`
	Subscription Subscription `json:"subscription"`
}

type Subscription struct {
	Depth    int    `json:"depth,omitempty"`
	Interval int    `json:"interval,omitempty"`
	Name     string `json:"name"`
	Token    string `json:"token,omitempty"`
}

type TickerInformation struct {
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

func NewKrakenClient() KrakenClient {
	url := url.URL{Scheme: "wss", Host: krakenWSUrl}
	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return KrakenClient{
		conn: c,
	}
}

func (k *KrakenClient) Subscribe(subscribeEvent SubscribeEvent) {
	event, err := json.Marshal(subscribeEvent)
	log.Println(string(event))
	if err != nil {
		log.Fatal("could not marshal subscribe event:", err)
	}
	k.conn.WriteMessage(websocket.TextMessage, event)
	input := make(chan string)
	go func(input chan string) {
		for {
			_, message, err := k.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			input <- string(message)
		}
	}(input)

	for i := 0; i < 5; i++ {
		log.Printf("recv: %s", <-input)
		time.Sleep(2 * time.Second)
	}
}

func (k *KrakenClient) Close() {
	k.conn.Close()
}
