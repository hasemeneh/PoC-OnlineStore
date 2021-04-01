package definitions

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
)

type ClaimDefinition interface {
	ClaimProduct(ctx context.Context, req *models.ClaimRequest) (*models.ClaimResponse, error)
	ValidateProductStock(ctx context.Context, req *models.ProductRequest) (*models.ProductResponse, error)
}
