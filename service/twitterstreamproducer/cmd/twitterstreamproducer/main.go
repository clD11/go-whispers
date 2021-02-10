package main

import (
	"fmt"
	. "github.com/clD11/go-whispers/service/twitterstreamproducer/internal"
	"time"
)

func main() {
	fmt.Println("Starting Twitter Stream Producer")

	twitterStreamWrapperConfig := TwitterStreamWrapperConfig{
		APIKey:            "test",
		APISecretKey:      "test",
		AccessToken:       "test",
		AccessTokenSecret: "test",
	}

	twitterStreamWrapperFilter := TwitterStreamWrapperFilter{
		Track: []string{"ETH,BTC"},
	}

	twitterStreamWrapper := NewTwitterStreamWrapper(twitterStreamWrapperConfig, twitterStreamWrapperFilter)
	twitterStreamWrapper.Start()

	time.Sleep(30 * time.Second)

	fmt.Println("Stopping Twitter Stream Producer")

	twitterStreamWrapper.Stop()
}
