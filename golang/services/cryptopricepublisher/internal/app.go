package internal

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Application struct {
	server   *http.Server
	router   *mux.Router
	shutdown chan bool
}

func NewApplication() Application {
	return Application{}
}

func (a *Application) Start(host string) {
	log.Println("starting server...")

	a.router = mux.NewRouter()
	a.router.HandleFunc("/api/v1/app/stop", a.Stop).Methods(http.MethodGet)
	a.server = &http.Server{Addr: host, Handler: a.router}
	a.shutdown = make(chan bool)

	go func() {
		a.shutdown <- true
		log.Println("shutting down server...")
		close(a.shutdown)
		a.server.SetKeepAlivesEnabled(false)
		a.server.Close()
	}()

	a.server.ListenAndServe()
}

func (a *Application) Stop(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server shutdown"))
	w.WriteHeader(200)
	<-a.shutdown
}
