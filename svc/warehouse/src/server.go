package main

import (
	"github.com/hasemeneh/PoC-OnlineStore/helper/webservice"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/config"
	http_backdoor "github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/handler/http/backdoor"
	http_public "github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/handler/http/public"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/service"
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
	ws.Run()
}
