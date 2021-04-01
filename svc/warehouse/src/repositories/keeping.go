package repositories

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
	"github.com/jmoiron/sqlx"
)

type KeepingRepo interface {
	StartTx(ctx context.Context) (*sqlx.Tx, error)
	GetProductByProductID(ctx context.Context, productID int64) (*models.KeepingModel, error)
	GetProductByProductIDWithLock(ctx context.Context, dbtx *sqlx.Tx, productID int64) (*models.KeepingModel, error)
	UpdateStockQuantityByProductID(ctx context.Context, dbtx *sqlx.Tx, productID int64, quantity int64) error
}
