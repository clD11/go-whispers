package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal"
)

var sut internal.App
var server *http.Server

func TestMain(m *testing.M) {
	sut = internal.App{}
	sut.Initialize()
	server = &http.Server{Addr: ":9807", Handler: sut.Router}

	code := m.Run()
	os.Exit(code)
}
