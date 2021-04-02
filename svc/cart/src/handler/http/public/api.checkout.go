package public

import (
	"net/http"
	"strconv"

	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
)

func (p *Public) HandleCheckout(r *http.Request) response.HttpResponse {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		return response.NewJsonResponse().SetError(response.ErrorBadRequest).SetMessage("Invalid Request")
	}
	resp, err := p.Service.Usecase.Checkout.CheckoutOrder(r.Context(), userID)
	if err != nil {
		return response.NewJsonResponse().SetError(err).SetMessage(err.Error())
	}
	return response.NewJsonResponse().SetData(resp)

}
