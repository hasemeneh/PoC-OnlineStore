package public

import (
	"net/http"
	"strconv"

	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
)

func (a *Public) HandleValidateStock(r *http.Request) response.HttpResponse {
	productID, err := strconv.ParseInt(r.URL.Query().Get("product_id"), 10, 64)
	if err != nil {
		return response.NewJsonResponse().SetError(response.ErrorBadRequest).SetMessage("Invalid Request")
	}
	demand, err := strconv.ParseInt(r.URL.Query().Get("demand"), 10, 64)
	if err != nil {
		return response.NewJsonResponse().SetError(response.ErrorBadRequest).SetMessage("Invalid Request")
	}
	req := &models.ProductRequest{
		ProductID: productID,
		Demand:    demand,
	}
	resp, err := a.Service.Usecase.Claim.ValidateProductStock(r.Context(), req)
	if err != nil {
		return response.NewJsonResponse().SetError(err)
	}
	return response.NewJsonResponse().SetData(resp)
}
