package public

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
)

func (p *Public) HandleGetActiveCart(r *http.Request) response.HttpResponse {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		return response.NewJsonResponse().SetError(response.ErrorBadRequest).SetMessage("Invalid Request")
	}
	resp, err := p.Service.Usecase.Checkout.GetCurrentUserCart(r.Context(), userID)
	if err != nil {
		return response.NewJsonResponse().SetError(err).SetMessage(err.Error())
	}
	return response.NewJsonResponse().SetData(resp)
}

func (p *Public) HandleAddToCart(r *http.Request) response.HttpResponse {
	request := models.AddToCartRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusBadRequest))
	}
	resp, err := p.Service.Usecase.Checkout.AddToCart(r.Context(), &request)
	if err != nil {
		return response.NewJsonResponse().SetError(err).SetMessage(err.Error())
	}
	return response.NewJsonResponse().SetData(resp)
}
