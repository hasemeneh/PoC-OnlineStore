package definitions

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
)

type CheckoutDefinition interface {
	GetUserCheckouts(ctx context.Context, userID int64) ([]models.UserRecordedCheckout, error)
	AddToCart(ctx context.Context, req *models.AddToCartRequest) (*models.AddToCartResponse, error)
	GetCurrentUserCart(ctx context.Context, userID int64) (*models.AddToCartResponse, error)
	CheckoutOrder(ctx context.Context, userID int64) (*models.CheckoutResponse, error)
}
