package main

import (
	"encoding/json"
	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal/handler"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal"
)

var sut internal.Application
var server *http.Server

func TestMain(m *testing.M) {
	sut = internal.Application{}
	sut.Initialize()
	server = &http.Server{Addr: ":9807", Handler: sut.Router}

	code := m.Run()
	os.Exit(code)
}

func TestGetStuff(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	rw := httptest.NewRecorder()

	server.Handler.ServeHTTP(rw, r)

	var actual handler.Response
	err := json.NewDecoder(rw.Body).Decode(&actual)

	assert.NoError(t, err)
	assert.Equal(t, "cryptocurrency price publisher v1", actual.Data)
}

func TestPostStuff(t *testing.T) {
	assert.Equal(t, true, true)
}
