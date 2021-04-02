package user

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
)

func (u *userUsecase) GetOrderByUserID(ctx context.Context, userID int64) ([]models.UserOrder, error) {
	resp := make([]models.UserOrder, 0)
	orders, err := u.OrderRepo.GetOrderByUserID(ctx, userID)
	if err != nil {
		return resp, err
	}

	items, err := u.ItemRepo.GetItemByUserID(ctx, userID)
	if err != nil {
		return resp, err
	}
	for _, order := range orders {
		item := items[order.ID]
		resp = append(resp, models.UserOrder{
			OrderModel: order,
			Items:      item,
		})
	}
	return resp, nil
}
