package control

import (
	"context"
	"log"
	"net/http"

	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
)

func (c *controlUsecase) GetStockByProductID(ctx context.Context, productID int64) (*models.StockResponse, error) {
	resp := &models.StockResponse{}
	var err error
	resp.KeepingModel, err = c.KeepingRepo.GetProductByProductID(ctx, productID)
	if err != nil {
		log.Println("[GetStockByProductID] Error On Getting stock keeping because of ", err)
		return resp, response.NewResponseError(err.Error(), http.StatusBadRequest)
	}
	resp.History, err = c.HistoryRepo.GetProductByProductID(ctx, productID)
	if err != nil {
		log.Println("[GetStockByProductID] Error On Getting stock history because of ", err)
		return resp, err
	}
	return resp, err
}
