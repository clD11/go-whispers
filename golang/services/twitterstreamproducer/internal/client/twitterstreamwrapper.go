package client

//
//import (
//	"fmt"
//	"github.com/dghubble/go-twitter/twitter"
//	"github.com/dghubble/oauth1"
//	"log"
//)
//
//type TwitterStreamWrapperConfig struct {
//	APIKey            string
//	APISecretKey      string
//	AccessToken       string
//	AccessTokenSecret string
//}
//
//type TwitterStreamWrapperFilter struct {
//	Track []string
//}
//
//type TwitterStreamWrapper struct {
//	twitterStream *twitter.Stream
//}
//
//func NewTwitterStreamWrapper(twitterStreamWrapperConfig TwitterStreamWrapperConfig,
//	twitterStreamWrapperFilter TwitterStreamWrapperFilter) TwitterStreamWrapper {
//
//	config := oauth1.NewConfig(twitterStreamWrapperConfig.APIKey, twitterStreamWrapperConfig.APISecretKey)
//	token := oauth1.NewToken(twitterStreamWrapperConfig.AccessToken, twitterStreamWrapperConfig.AccessTokenSecret)
//	httpClient := config.Client(oauth1.NoContext, token)
//
//	client := twitter.NewClient(httpClient)
//
//	filterParams := &twitter.StreamFilterParams{
//		Track:         twitterStreamWrapperFilter.Track,
//		StallWarnings: twitter.Bool(false),
//	}
//
//	twitterStream, err := client.Streams.Filter(filterParams)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return TwitterStreamWrapper{
//		twitterStream: twitterStream,
//	}
//}
//
//func (t *TwitterStreamWrapper) Start() {
//	demux := twitter.NewSwitchDemux()
//	demux.Tweet = func(tweet *twitter.Tweet) {
//		fmt.Println(tweet.Text)
//	}
//	go demux.HandleChan(t.twitterStream.Messages)
//}
//
//func (t *TwitterStreamWrapper) Stop() {
//	t.twitterStream.Stop()
//}
