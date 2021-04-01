package grpc

import (
	pb "github.com/hasemeneh/PoC-OnlineStore/svc/grpc/warehouse/protos"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"

	"context"
)

func (s *server) ValidateStock(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	resp, err := s.Service.Usecase.Claim.ValidateProductStock(ctx, &models.ProductRequest{
		ProductID: req.ProductID,
		Demand:    req.Demand,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ProductResponse{
		ProductID: resp.ProductID,
		Demand:    resp.Demand,
		Status:    int64(resp.Status),
	}, nil

}
func (s *server) ClaimProduct(ctx context.Context, req *pb.ClaimRequest) (*pb.ClaimResponse, error) {
	resp, err := s.Service.Usecase.Claim.ClaimProduct(ctx, translacteClaimProductRequest(req))
	if err != nil {
		return nil, err
	}
	return translacteClaimProductResponse(resp), nil
}
func translacteClaimProductRequest(req *pb.ClaimRequest) *models.ClaimRequest {
	resp := &models.ClaimRequest{
		Products: make([]models.ProductRequest, 0),
	}
	for _, product := range req.Products {
		resp.Products = append(resp.Products, models.ProductRequest{
			Demand:    product.Demand,
			ProductID: product.ProductID,
		})
	}
	return resp
}
func translacteClaimProductResponse(req *models.ClaimResponse) *pb.ClaimResponse {
	resp := &pb.ClaimResponse{
		Products: []*pb.ProductResponse{},
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
