package grpc

import (
	"context"

	pb "github.com/hasemeneh/PoC-OnlineStore/svc/grpc/order/protos"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
)

func (s *server) Checkout(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
	resp, err := s.Service.Usecase.Transaction.CreateOrder(ctx, translacteOrderProductRequest(in))
	if err != nil {
		return nil, err
	}
	return translacteOrderProductResponse(resp), nil
}
func translacteOrderProductRequest(req *pb.OrderRequest) *models.OrderCreationRequest {
	resp := &models.OrderCreationRequest{
		Products: make([]models.ProductRequest, 0),
		CartID:   req.CartID,
		UserID:   req.UserID,
	}
	for _, product := range req.Products {
		resp.Products = append(resp.Products, models.ProductRequest{
			Demand:    product.Demand,
			ProductID: product.ProductID,
		})
	}
	return resp
}
func translacteOrderProductResponse(req *models.OrderCreationResponse) *pb.OrderResponse {
	resp := &pb.OrderResponse{
		IsSuccess: req.IsSuccess,
		Products:  []*pb.ProductResponse{},
	}

	for _, product := range req.Products {
		resp.Products = append(resp.Products, &pb.ProductResponse{
			ProductID: product.ProductID,
			Demand:    product.Demand,
			Status:    int64(product.Status),
		})
	}
	return resp
}
