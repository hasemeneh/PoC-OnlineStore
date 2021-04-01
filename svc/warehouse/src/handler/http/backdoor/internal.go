package backdoor

import (
	"net/http"

	"github.com/hasemeneh/PoC-OnlineStore/helper/httphandler"
	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/service"
	"github.com/julienschmidt/httprouter"
)

type Internal struct {
	Service *service.Service
}

func NewHandler(Service *service.Service) *Internal {
	return &Internal{
		Service: Service,
	}
}
func (i *Internal) Register(r *httprouter.Router) {
	apiHttpHandlers := httphandler.New("/internal", r)
	apiHttpHandlers.GET("/ping", i.PING)

}
func (i *Internal) PING(r *http.Request) response.HttpResponse {
	return response.NewJsonResponse().SetMessage("Pong").SetData("Pung")
}
