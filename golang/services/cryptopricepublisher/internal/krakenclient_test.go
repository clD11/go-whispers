package internal

import (
	"fmt"
	"testing"
)

func TestGetTickerInformation(t *testing.T) {
	fmt.Println("Ending")
}

func TestSubscribe(t *testing.T) {
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
