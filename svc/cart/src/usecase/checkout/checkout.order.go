package checkout

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/constants"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
)

func (c *checkoutUsecase) CheckoutOrder(ctx context.Context, userID int64) (*models.CheckoutResponse, error) {

	resp := &models.CheckoutResponse{}
	cartCache, err := c.GetCurrentUserCart(ctx, userID)
	if err != nil {
		return nil, err
	}
	if len(cartCache.Products) < 1 {
		return nil, response.NewResponseError("empty cart", http.StatusUnprocessableEntity)
	}
	dbtx, err := c.CartRepo.StartTx(ctx)
	if err != nil {
		return resp, err
	}
	defer dbtx.Rollback()
	cartID, err := c.CartRepo.InsertCart(ctx, dbtx, models.CartModel{
		UserID: userID,
	})
	if err != nil {
		return resp, err
	}
	for _, product := range cartCache.Products {
		_, err = c.ItemRepo.InsertItem(ctx, dbtx, models.ItemModels{
			CartID:    cartID,
			ProductID: product.ProductID,
			Quantity:  product.Demand,
		})
		if err != nil {
			return resp, err
		}
	}
	orderResp, err := c.Order.Checkout(ctx, &models.OrderCreationRequest{
		CartID:   cartID,
		UserID:   userID,
		Products: cartCache.Products,
	})
	if err != nil {
		return resp, err
	}
	for _, orderProduct := range orderResp.Products {
		if orderProduct.Status != constants.ClaimStatusOK {
			return resp, response.NewResponseError(fmt.Sprintf("Order not created because productID %d is insufficient", orderProduct.ProductID), http.StatusUnprocessableEntity)
		}
	}
	if !orderResp.IsSuccess {
		return resp, response.NewResponseError("Order not created", http.StatusUnprocessableEntity)
	}
	resp.Products = orderResp.Products
	err = c.CartRepo.DestroyCart(ctx, userID)
	if err != nil {
		log.Println("[CheckoutOrder] Failed to destroy Cart Session")
	}
	return resp, nil
}
