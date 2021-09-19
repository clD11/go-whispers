package main

import (
	"encoding/json"
	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

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

func TestHealth(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	rw := httptest.NewRecorder()

	server.Handler.ServeHTTP(rw, request)

	var actual handler.HealthResponse
	json.NewDecoder(rw.Body).Decode(&actual)

	expected := handler.HealthResponse{
		Service:   "crypto price publisher",
		Status:    "healthy",
		Timestamp: time.Now().String(),
	}

	assert.Equal(t, expected.Service, actual.Service)
	assert.Equal(t, expected.Status, actual.Status)
}

func TestWebsocket(t *testing.T) {

}
