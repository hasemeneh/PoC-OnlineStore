package order

import (
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/constants"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
	order_GRPC "github.com/hasemeneh/PoC-OnlineStore/svc/grpc/order/protos"

	"context"
)

func (o *OrderInterface) Checkout(ctx context.Context, in *models.OrderCreationRequest) (*models.OrderCreationResponse, error) {
	resp, err := o.client.Checkout(ctx, translateCheckoutOrderRequest(in))
	if err != nil {
		return nil, err
	}
	return translateCheckoutOrderResponse(resp), nil
}
func translateCheckoutOrderRequest(req *models.OrderCreationRequest) *order_GRPC.OrderRequest {
	resp := &order_GRPC.OrderRequest{
		Products: make([]*order_GRPC.ProductRequest, 0),
		CartID:   req.CartID,
		UserID:   req.UserID,
	}
	for _, product := range req.Products {
		resp.Products = append(resp.Products, &order_GRPC.ProductRequest{
			ProductID: product.ProductID,
			Demand:    product.Demand,
		})
	}
	return resp
}

func translateCheckoutOrderResponse(req *order_GRPC.OrderResponse) *models.OrderCreationResponse {
	resp := &models.OrderCreationResponse{
		Products:  make([]models.ProductResponse, 0),
		IsSuccess: req.IsSuccess,
	}
	for _, product := range req.Products {
		resp.Products = append(resp.Products, models.ProductResponse{
			ProductRequest: models.ProductRequest{
				Demand:    product.Demand,
				ProductID: product.ProductID,
			},
			Status: constants.ClaimStatus(product.Status),
		})
	}
	return resp
}
