package repositories

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
	"github.com/jmoiron/sqlx"
)

type ItemRepo interface {
	InsertItem(ctx context.Context, dbtx *sqlx.Tx, req models.ItemModels) (int64, error)
	StartTx(ctx context.Context) (*sqlx.Tx, error)
	GetItemByUserID(ctx context.Context, UserID int64) (map[int64][]models.ItemModels, error)
}
