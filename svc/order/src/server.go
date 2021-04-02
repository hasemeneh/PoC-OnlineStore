package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hasemeneh/PoC-OnlineStore/helper/webservice"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/config"
	grpc_server "github.com/hasemeneh/PoC-OnlineStore/svc/order/src/handler/grpc"
	http_backdoor "github.com/hasemeneh/PoC-OnlineStore/svc/order/src/handler/http/backdoor"
	http_public "github.com/hasemeneh/PoC-OnlineStore/svc/order/src/handler/http/public"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/service"
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
	grpcServer := grpc_server.New(&grpc_server.Option{
		Service:  serviceObj,
		GrpcPort: cfg.GRPCPort,
	})
	go grpcServer.Serve()
	go ws.Run()
	select {
	case s := <-terminateSignal():
		grpcServer.CatchSignal(s)
		log.Println("Exiting gracefully...")
	}
}

func terminateSignal() chan os.Signal {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	return term
}
