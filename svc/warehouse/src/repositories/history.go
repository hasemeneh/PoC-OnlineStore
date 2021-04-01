package repositories

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
	"github.com/jmoiron/sqlx"
)

type HistoryRepo interface {
	StartTx(ctx context.Context) (*sqlx.Tx, error)
	GetProductByProductID(ctx context.Context, productID int64) ([]models.StockHistoryModel, error)
	InsertHistory(ctx context.Context, dbtx *sqlx.Tx, req models.StockHistoryModel) error
}
