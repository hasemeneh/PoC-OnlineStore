package webservice

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type WebService interface {
	Run()
}
type logger func(v ...interface{})

type WebRegistrator interface {
	Register(router *httprouter.Router)
}
type WebHandler struct {
	router       *httprouter.Router
	registrators []WebRegistrator
	port         string
	fatalln      logger
}

func NewWebService(port string, registrators ...WebRegistrator) WebService {
	return &WebHandler{
		router:       httprouter.New(),
		port:         port,
		registrators: registrators,
		fatalln:      log.Fatalln,
	}
}

func (w *WebHandler) Run() {
	if w.port == "" {
		w.fatalln("Empty port defined")
	}
	for _, registrator := range w.registrators {
		registrator.Register(w.router)
	}
	log.Println("Server running on localhost" + w.port)
	log.Fatalln(http.ListenAndServe(w.port, w.router))
}
