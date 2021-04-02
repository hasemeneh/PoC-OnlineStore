package checkout

import (
	"context"
	"net/http"

	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/constants"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
)

func (c *checkoutUsecase) AddToCart(ctx context.Context, req *models.AddToCartRequest) (*models.AddToCartResponse, error) {
	warehouseResp, err := c.Warehouse.ValidateStock(ctx, &req.ProductRequest)
	if err != nil {
		return nil, err
	}
	if warehouseResp.Status != constants.ClaimStatusOK {
		return nil, response.NewResponseError("Insufficient Stock", http.StatusUnprocessableEntity)
	}

	return c.CartRepo.AddToCart(ctx, req)
}

func (c *checkoutUsecase) GetCurrentUserCart(ctx context.Context, userID int64) (*models.AddToCartResponse, error) {

	return c.CartRepo.GetUserCart(ctx, userID)
}
