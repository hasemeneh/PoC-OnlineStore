package public

import (
	"net/http"

	"github.com/hasemeneh/PoC-OnlineStore/helper/httphandler"
	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/service"
	"github.com/julienschmidt/httprouter"
)

type Public struct {
	Service *service.Service
}

func NewHandler(Service *service.Service) *Public {
	return &Public{
		Service: Service,
	}
}
func (p *Public) Register(r *httprouter.Router) {
	apiHttpHandlers := httphandler.New("/api", r)
	apiHttpHandlers.GET("/ping", p.PING)
	apiHttpHandlers.GET("/cart", p.HandleGetActiveCart)
	apiHttpHandlers.PUT("/cart", p.HandleAddToCart)
	apiHttpHandlers.POST("/checkout", p.HandleCheckout)
}
func (p *Public) PING(r *http.Request) response.HttpResponse {
	return response.NewJsonResponse().SetMessage("Pong").SetData("Pung")
}
