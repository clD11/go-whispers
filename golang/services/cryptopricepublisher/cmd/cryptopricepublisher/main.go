package main

import (
	"fmt"
	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal"
)

func main() {
	internal.GetTickerInformation()
	fmt.Println("End Client")
}
