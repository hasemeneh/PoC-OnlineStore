package claim

import (
	"context"
	"errors"

	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/constants"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
)

func (cu *claimUsecase) ClaimProduct(ctx context.Context, req *models.ClaimRequest) (*models.ClaimResponse, error) {
	resp := &models.ClaimResponse{}
	resp.Products = make([]models.ProductResponse, 0)
	dbtx, err := cu.KeepingRepo.StartTx(ctx)
	if err != nil {
		return resp, err
	}
	defer dbtx.Rollback()
	for _, product := range req.Products {
		keeping, err := cu.KeepingRepo.GetProductByProductIDWithLock(ctx, dbtx, product.ProductID)
		if err != nil {
			return resp, err
		}
		status := constants.ClaimStatusOutOfStock
		if keeping.Quantity >= product.Demand {
			status = constants.ClaimStatusOK
			err = cu.KeepingRepo.UpdateStockQuantityByProductID(ctx, dbtx, keeping.ProductID, keeping.Quantity-product.Demand)
			if err != nil {
				return resp, err
			}
			err = cu.HistoryRepo.InsertHistory(ctx, dbtx, models.StockHistoryModel{
				Amount:         product.Demand,
				StockKeepingID: keeping.ID,
				Type:           constants.HistoryTypeOutcoming,
			})
			if err != nil {
				return resp, err
			}
		}
		resp.Products = append(resp.Products,
			models.ProductResponse{
				ProductRequest: &product,
				Status:         status,
			},
		)
	}
	return resp, dbtx.Commit()
}

func (cu *claimUsecase) ValidateProductStock(ctx context.Context, req *models.ProductRequest) (*models.ProductResponse, error) {
	resp := &models.ProductResponse{
		ProductRequest: req,
		Status:         constants.ClaimStatusOutOfStock,
	}
	if req == nil {
		return resp, errors.New("Bad Request")
	}
	stock, err := cu.KeepingRepo.GetProductByProductID(ctx, req.ProductID)
	if err != nil {
		return resp, err
	}
	if stock.Quantity >= req.Demand {
		resp.Status = constants.ClaimStatusOK
	}
	return resp, nil
}
