package internal

import (
	"fmt"
	"testing"
)

func TestGetTickerInformation(t *testing.T) {
	fmt.Println("Ending")
}

func TestSubscribe(t *testing.T) {
	go startProcess()
	subscribeEvent := SubscribeEvent{
		Event: "subscribe",
		Pair:  []string{"XBT/USD"},
		Subscription: Subscription{
			Name: "ticker",
		},
	}
	krakenClient := NewKrakenClient()
	krakenClient.Subscribe(subscribeEvent)
}

func startProcess() {
	subscribeEvent := SubscribeEvent{
		Event: "subscribe",
		Pair:  []string{"XBT/USD"},
		Subscription: Subscription{
			Name: "ticker",
		},
	}
	krakenClient := NewKrakenClient()
	krakenClient.Subscribe(subscribeEvent)
}
