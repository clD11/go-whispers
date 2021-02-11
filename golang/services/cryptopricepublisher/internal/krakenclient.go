package internal

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

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

type KrakenResponseWrapper struct {
	result string   `json:"result"`
	error  []string `json:"error"`
}

type TickerEvent struct {
	event        string       `json:"event"`
	pair         []string     `json:"pair"`
	subscription Subscription `json:"subscription"`
}

type Subscription struct {
	name string `json:"name"`
}

const (
	krakenWebsocket = "ws.kraken.com"
	event           = "{\"event\":\"subscribe\",\"pair\":[\"XBT/USD\"],\"subscription\":{\"name\":\"ticker\"}}"
)

// WIP - add routines and channels
func GetTickerInformation() {
	url := url.URL{Scheme: "wss", Host: krakenWebsocket}
	log.Printf("connecting to %s", url.String())

	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	c.WriteMessage(websocket.TextMessage, []byte(event))

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}
