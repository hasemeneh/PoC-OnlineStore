package transaction

import (
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/interface/warehouse"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/repositories"
)

type transactionUsecase struct {
	ItemRepo  repositories.ItemRepo
	OrderRepo repositories.OrderRepo
	Warehouse warehouse.WarehouseManager
}
type Option struct {
	ItemRepo  repositories.ItemRepo
	OrderRepo repositories.OrderRepo
	Warehouse warehouse.WarehouseManager
}

func New(o *Option) *transactionUsecase {
	return &transactionUsecase{
		ItemRepo:  o.ItemRepo,
		OrderRepo: o.OrderRepo,
		Warehouse: o.Warehouse,
	}
}
