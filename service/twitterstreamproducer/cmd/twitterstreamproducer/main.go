package main

import (
	. "github.com/clD11/go-whispers/service/twitterstreamproducer/internal"
	"time"
)

func main() {
	twitterStreamWrapperConfig := TwitterStreamWrapperConfig{
		APIKey:            "ajpcpacmp",
		APISecretKey:      "nosdnvosdnvonsdovnosd",
		AccessToken:       "snidviosdnvionsdfviondfiovndfiovnosdfnvoi",
		AccessTokenSecret: "vcjsopvmpvmpsdmvpsdmvpdspovmopsdmvopmsdopmvpmsdpvm",
	}
	twitterStreamWrapperFilter := TwitterStreamWrapperFilter{
		Track: []string{"ETH,BTC"},
	}

	twitterStreamWrapper := NewTwitterStreamWrapper(twitterStreamWrapperConfig, twitterStreamWrapperFilter)
	twitterStreamWrapper.Start()

	time.Sleep(60 * time.Second)

	twitterStreamWrapper.Stop()
}
