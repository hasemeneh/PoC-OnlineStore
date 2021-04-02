package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hasemeneh/PoC-OnlineStore/helper/webservice"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/config"
	http_backdoor "github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/handler/http/backdoor"
	http_public "github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/handler/http/public"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/service"
)

func main() {
	cfg := config.New().Read()
	serviceObj := service.New(cfg)
	publicHttpHandler := http_public.NewHandler(serviceObj)
	internalHttpHandler := http_backdoor.NewHandler(serviceObj)
	ws := webservice.NewWebService(
		cfg.RunningPort,
		publicHttpHandler,
		internalHttpHandler,
	)
	go ws.Run()
	select {
	case s := <-terminateSignal():
		log.Println("Exiting gracefully...", s)
	}
}

func terminateSignal() chan os.Signal {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	return term
}
