package checkout

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
)

func (c *checkoutUsecase) GetUserCheckouts(ctx context.Context, userID int64) ([]models.UserRecordedCheckout, error) {
	resp := make([]models.UserRecordedCheckout, 0)
	carts, err := c.CartRepo.GetRecordedCartByUserID(ctx, userID)
	if err != nil {
		return resp, err
	}
	items, err := c.ItemRepo.GetItemByUserID(ctx, userID)
	if err != nil {
		return resp, err
	}
	for _, cart := range carts {
		resp = append(resp, models.UserRecordedCheckout{
			CartModel: cart,
			Item:      items[cart.ID],
		})

	}
	return resp, nil

}
