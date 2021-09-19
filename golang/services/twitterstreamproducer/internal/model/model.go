package model

type WhisperMessage struct {
	transactionId string   `json:"transaction_id"`
	data          string   `json:"data"`
	error         []string `json:"error"`
}
