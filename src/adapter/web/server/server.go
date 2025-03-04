package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/adapter/web/handler"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/app"
)

type WebServer struct {
	Service app.IProductService
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Main() {
	router := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handler.MakeProductHandlers(router, n, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		Addr:    ":9000",
		Handler: http.DefaultServeMux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 	 10 * time.Second,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	
	if err != nil {
		log.Fatal(err)
	}
}
