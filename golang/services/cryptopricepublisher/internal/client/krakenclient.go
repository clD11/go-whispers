package client

//
//import (
//	"encoding/json"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/url"
//	"time"
//)
//
//const (
//	krakenWSUrl = "ws.kraken.com"
//)
//
//type KrakenClient struct {
//	conn *websocket.Conn
//}
//
//type SubscribeEvent struct {
//	Event        string       `json:"event"`
//	Reqid        int          `json:"reqid,omitempty"`
//	Pair         []string     `json:"pair"`
//	Subscription Subscription `json:"subscription"`
//}
//
//type Subscription struct {
//	Depth    int    `json:"depth,omitempty"`
//	Interval int    `json:"interval,omitempty"`
//	Name     string `json:"name"`
//	Token    string `json:"token,omitempty"`
//}
//
//type TickerInformation struct {
//	ask                        []string `json:"a"`
//	bid                        []string `json:"b"`
//	lastTradeClosed            []string `json:"c"`
//	volume                     []string `json:"v"`
//	volumeWeightedAveragePrice []string `json:"p"`
//	numberOfTrades             []string `json:"t"`
//	low                        []string `json:"l"`
//	high                       []string `json:"h"`
//	todayOpenPrice             string   `json:"o"`
//}
//
//func NewKrakenClient() KrakenClient {
//	url := url.URL{Scheme: "wss", Host: krakenWSUrl}
//	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
//	if err != nil {
//		log.Fatal("dial:", err)
//	}
//	return KrakenClient{
//		conn: c,
//	}
//}
//
//var count = 0
//
//func (k *KrakenClient) Subscribe(subscribeEvent SubscribeEvent) {
//	count += 1
//	id := count
//	event, err := json.Marshal(subscribeEvent)
//	log.Println(string(event))
//	if err != nil {
//		log.Fatal("could not marshal subscribe event:", err)
//	}
//	k.conn.WriteMessage(websocket.TextMessage, event)
//	input := make(chan string)
//	go func(input chan string) {
//		for {
//			_, message, err := k.conn.ReadMessage()
//			if err != nil {
//				log.Println("read:", err)
//				return
//			}
//			log.Printf("process_id: %d %s\n", id, string(message))
//			//input <- string(message)
//		}
//	}(input)
//
//	for i := 0; i < 5; i++ {
//		log.Printf("waiting: %d ses", 5-i)
//		time.Sleep(5 * time.Second)
//	}
//	log.Println("*** END PROCESS ***")
//}
//
//func (k *KrakenClient) Close() {
//	k.conn.Close()
//}
