package backdoor

import (
	"net/http"
	"strconv"

	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
)

func (b *Internal) HandleGetStock(r *http.Request) response.HttpResponse {

	productID, err := strconv.ParseInt(r.URL.Query().Get("product_id"), 10, 64)
	if err != nil {
		return response.NewJsonResponse().SetError(response.ErrorBadRequest).SetMessage("Invalid Request")
	}
	resp, err := b.Service.Usecase.Control.GetStockByProductID(r.Context(), productID)
	if err != nil {
		return response.NewJsonResponse().SetError(err)
	}
	return response.NewJsonResponse().SetData(resp)

}
