package warehouse

import (
	"context"

	pb "github.com/hasemeneh/PoC-OnlineStore/svc/grpc/warehouse/protos"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/constants"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
)

func (w *warehouseInterface) ClaimProduct(ctx context.Context, req *models.ClaimRequest) (*models.ClaimResponse, error) {
	resp, err := w.client.ClaimProduct(ctx, translacteClaimProductRequest(req))
	if err != nil {
		return nil, err
	}
	return translacteClaimProductResponse(resp), nil
}

func translacteClaimProductRequest(req *models.ClaimRequest) *pb.ClaimRequest {
	resp := &pb.ClaimRequest{
		Products: make([]*pb.ProductRequest, 0),
		OrderID:  req.OrderID,
		CartID:   req.CartID,
	}
	for _, product := range req.Products {
		resp.Products = append(resp.Products, &pb.ProductRequest{
			Demand:    product.Demand,
			ProductID: product.ProductID,
		})
	}
	return resp
}
func translacteClaimProductResponse(req *pb.ClaimResponse) *models.ClaimResponse {
	resp := &models.ClaimResponse{
		Products: []models.ProductResponse{},
	}

	for _, product := range req.Products {
		resp.Products = append(resp.Products, models.ProductResponse{
			ProductRequest: models.ProductRequest{
				ProductID: product.ProductID,
				Demand:    product.Demand,
			},
			Status: constants.ClaimStatus(product.Status),
		})
	}
	return resp
}
