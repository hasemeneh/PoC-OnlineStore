package checkout

import (
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/interface/order"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/interface/warehouse"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/repositories"
)

type checkoutUsecase struct {
	ItemRepo  repositories.ItemRepo
	CartRepo  repositories.CartRepo
	Warehouse warehouse.WarehouseManager
	Order     order.OrderManager
}
type Option struct {
	ItemRepo  repositories.ItemRepo
	CartRepo  repositories.CartRepo
	Warehouse warehouse.WarehouseManager
	Order     order.OrderManager
}

func New(o *Option) *checkoutUsecase {
	return &checkoutUsecase{
		ItemRepo:  o.ItemRepo,
		CartRepo:  o.CartRepo,
		Warehouse: o.Warehouse,
		Order:     o.Order,
	}
}
