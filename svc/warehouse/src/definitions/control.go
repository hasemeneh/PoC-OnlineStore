package definitions

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
)

type ControlDefinition interface {
	GetStockByProductID(ctx context.Context, productID int64) (*models.StockResponse, error)
}
