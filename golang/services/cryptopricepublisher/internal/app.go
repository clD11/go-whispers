package internal

import (
	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	server   *http.Server
	router   *mux.Router
	shutdown chan os.Signal
}

func NewApplication() Application {
	return Application{}
}

func (a *Application) Start(host string) {
	log.Println("Starting Server...")

	a.router = mux.NewRouter()
	a.router.HandleFunc("/", handler.GetLanding).Methods(http.MethodGet)

	a.server = &http.Server{Addr: host, Handler: a.router}

	a.shutdown = make(chan os.Signal)
	signal.Notify(a.shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-a.shutdown
		log.Println("Server Shutdown")
		a.server.SetKeepAlivesEnabled(false)
		a.server.Close()
		close(a.shutdown)
	}()

	log.Println("Listening")

	a.server.ListenAndServe()
}
