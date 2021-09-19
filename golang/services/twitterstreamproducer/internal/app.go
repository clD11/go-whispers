package internal

import (
	"github.com/clD11/go-whispers/golang/services/twitterstreamproducer/internal/handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	server   *http.Server
	Router   *mux.Router
	shutdown chan os.Signal
}

func (a *App) Health(w http.ResponseWriter, r *http.Request) {
	handler.Health(w, r)
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/health", a.Health).Methods(http.MethodGet)
}

func (a App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
	//a.server = &http.Server{Addr: host, Handler: a.Router}
	//
	//a.shutdown = make(chan os.Signal)
	//signal.Notify(a.shutdown, os.Interrupt, syscall.SIGTERM)
	//
	//go func() {
	//	<-a.shutdown
	//	a.server.SetKeepAlivesEnabled(false)
	//	a.server.Close()
	//	close(a.shutdown)
	//}()
	//
	//log.Println("listening")
	//
	//a.server.ListenAndServe()
}
