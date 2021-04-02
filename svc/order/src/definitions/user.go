package definitions

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
)

type UserDefinition interface {
	GetOrderByUserID(ctx context.Context, userID int64) ([]models.UserOrder, error)
}
