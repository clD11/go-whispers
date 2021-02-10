package internal

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

const krakenWebsocket = "ws://ws.kraken.com"
