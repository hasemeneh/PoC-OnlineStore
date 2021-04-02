package transaction

import (
	"context"
	"errors"

	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/constants"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
)

func (t *transactionUsecase) CreateOrder(ctx context.Context, req *models.OrderCreationRequest) (*models.OrderCreationResponse, error) {
	resp := &models.OrderCreationResponse{}
	dbtx, err := t.OrderRepo.StartTx(ctx)
	if err != nil {
		return resp, err
	}
	defer dbtx.Rollback()
	orderID, err := t.OrderRepo.InsertOrder(ctx, dbtx, models.OrderModel{
		UserID: req.UserID,
	})
	if err != nil {
		return resp, err
	}
	for _, product := range req.Products {
		_, err = t.ItemRepo.InsertItem(ctx, dbtx, models.ItemModels{
			OrderID:   orderID,
			ProductID: product.ProductID,
			Quantity:  product.Demand,
		})
		if err != nil {
			return resp, err
		}

	}
	warehouseResp, err := t.Warehouse.ClaimProduct(ctx, &models.ClaimRequest{
		CartID:   req.CartID,
		OrderID:  orderID,
		Products: req.Products,
	})
	if err != nil {
		return resp, err
	}
	if warehouseResp == nil {
		return resp, errors.New("Invalid Response from warehouse")
	}
	if len(warehouseResp.Products) != len(req.Products) {
		return resp, errors.New("Invalid Response from warehouse")
	}
	for _, warehouseProduct := range warehouseResp.Products {
		if warehouseProduct.Status != constants.ClaimStatusOK {
			return resp, errors.New("Product is Unavailable")
		}
		resp.Products = append(resp.Products, models.ProductResponse{
			ProductRequest: models.ProductRequest{
				ProductID: warehouseProduct.ProductID,
				Demand:    warehouseProduct.Demand,
			},
			Status: constants.ClaimStatusOK,
		})
	}
	resp.IsSuccess = true
	return resp, dbtx.Commit()
}
