package repositories

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
	"github.com/jmoiron/sqlx"
)

type OrderRepo interface {
	InsertOrder(ctx context.Context, dbtx *sqlx.Tx, req models.OrderModel) (int64, error)
	StartTx(ctx context.Context) (*sqlx.Tx, error)
	GetOrderByUserID(ctx context.Context, UserID int64) ([]models.OrderModel, error)
}
