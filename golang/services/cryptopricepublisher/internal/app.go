package internal

import (
	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal/handler"
	"github.com/clD11/go-whispers/golang/services/cryptopricepublisher/internal/websocket"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type App struct {
	server   *http.Server
	Router   *mux.Router
	wsPool   *websocket.Pool
	shutdown chan os.Signal
}

func (a *App) Health(rw http.ResponseWriter, r *http.Request) {
	handler.Health(rw, r)
}

func (a *App) SubscribePriceUpdate(rw http.ResponseWriter, r *http.Request) {
	handler.SubscribePriceUpdate(a.wsPool, rw, r)
}

func (a *App) Initialize() {
	a.InitializeRoots()
	a.InitializeWSPool()
}

func (a *App) InitializeWSPool() {
	a.wsPool = websocket.New()
}

func (a *App) InitializeRoots() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/health", a.Health).Methods(http.MethodGet)
	a.Router.HandleFunc("/subscribe-price-update", a.SubscribePriceUpdate).Methods(http.MethodGet)
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
	//	log.Println("Server Shutdown")
	//	a.server.SetKeepAlivesEnabled(false)
	//	a.server.Close()
	//	close(a.shutdown)
	//}()
	//
	//log.Println("Listening")
	//
	//a.server.ListenAndServe()
}
