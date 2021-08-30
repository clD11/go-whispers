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
	Router   *mux.Router
	shutdown chan os.Signal
}

func NewApplication() Application {
	return Application{}
}

func (a *Application) Initialize() {
	log.Println("Initializing Routes")
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", handler.GetLanding).Methods(http.MethodGet)
}

func (a Application) Run(host string) {
	a.server = &http.Server{Addr: host, Handler: a.Router}

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
