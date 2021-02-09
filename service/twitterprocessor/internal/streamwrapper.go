package internal

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type StreamWrapperConfig struct {
	apiKey            string
	apiSecretKey      string
	accessToken       string
	accessTokenSecret string
}

type Filter struct {
	track []string
}

func Start(streamWrapperConfig StreamWrapperConfig, filter Filter) {
	config := oauth1.NewConfig(streamWrapperConfig.apiKey, streamWrapperConfig.apiSecretKey)
	token := oauth1.NewToken(streamWrapperConfig.accessToken, streamWrapperConfig.accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = processor()

	filterParams := &twitter.StreamFilterParams{
		Track:         filter.track,
		StallWarnings: twitter.Bool(false),
	}

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	go demux.HandleChan(stream.Messages)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}

func processor() func(tweet *twitter.Tweet) {
	return func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}
}
