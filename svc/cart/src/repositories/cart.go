package repositories

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
	"github.com/jmoiron/sqlx"
)

type CartRepo interface {
	AddToCart(ctx context.Context, req *models.AddToCartRequest) (*models.AddToCartResponse, error)
	GetUserCart(ctx context.Context, userID int64) (*models.AddToCartResponse, error)
	InsertCart(ctx context.Context, dbtx *sqlx.Tx, req models.CartModel) (int64, error)
	StartTx(ctx context.Context) (*sqlx.Tx, error)
	GetRecordedCartByUserID(ctx context.Context, UserID int64) ([]models.CartModel, error)
	DestroyCart(ctx context.Context, userID int64) error
}
