package user

import (
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/repositories"
)

type userUsecase struct {
	ItemRepo  repositories.ItemRepo
	OrderRepo repositories.OrderRepo
}
type Option struct {
	ItemRepo  repositories.ItemRepo
	OrderRepo repositories.OrderRepo
}

func New(o *Option) *userUsecase {
	return &userUsecase{
		ItemRepo:  o.ItemRepo,
		OrderRepo: o.OrderRepo,
	}
}
