package definitions

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
)

type TransactionDefinitions interface {
	CreateOrder(ctx context.Context, req *models.OrderCreationRequest) (*models.OrderCreationResponse, error)
}
