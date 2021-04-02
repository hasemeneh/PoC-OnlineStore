package warehouse

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/constants"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
	pb "github.com/hasemeneh/PoC-OnlineStore/svc/grpc/warehouse/protos"
)

func (w *warehouseInterface) ValidateStock(ctx context.Context, req *models.ProductRequest) (*models.ProductResponse, error) {
	resp, err := w.client.ValidateStock(ctx, &pb.ProductRequest{
		ProductID: req.ProductID,
		Demand:    req.Demand,
	})
	if err != nil {
		return nil, err
	}
	return &models.ProductResponse{
		ProductRequest: models.ProductRequest{
			Demand:    req.Demand,
			ProductID: req.ProductID,
		},
		Status: constants.ClaimStatus(resp.Status),
	}, nil
}
