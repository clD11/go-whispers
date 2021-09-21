package client

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTickerInformation(t *testing.T) {
	fmt.Println("Ending")
}

func TestSubscribe(t *testing.T) {
	go startProcess()
	time.Sleep(time.Second * 60)
}

func startProcess() {
	subscribeEvent := SubscribeMessage{
		Event: "subscribe",
		Pair:  []string{"XBT/USD"},
		Subscription: Subscription{
			Name: "ticker",
		},
	}
	krakenClient := NewKrakenClient()
	krakenClient.SubscribePublic(subscribeEvent)
	krakenClient.ReadMessage()
}
