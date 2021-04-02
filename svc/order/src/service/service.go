package service

import (
	"log"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/definitions"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/domain/item"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/domain/order"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/interface/warehouse"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/usecase/transaction"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/usecase/user"
)

type Service struct {
	cfg       *models.MainConfig
	Interface struct {
		Warehouse warehouse.WarehouseManager
	}
	Usecase struct {
		Transaction definitions.TransactionDefinitions
		User        definitions.UserDefinition
	}
}

func New(cfg *models.MainConfig) *Service {
	db := database.New().Connect(cfg.DBConnectionString)
	itemModule := item.New(db)
	orderModule := order.New(db)
	warehouseInterface := warehouse.New(&warehouse.Option{
		Address: cfg.WarehouseAddress,
		Fatalln: log.Fatalln,
	})
	serviceObj := Service{
		cfg: cfg,
	}
	serviceObj.Usecase.Transaction = transaction.New(&transaction.Option{
		ItemRepo:  itemModule,
		OrderRepo: orderModule,
		Warehouse: warehouseInterface,
	})
	serviceObj.Usecase.User = user.New(&user.Option{
		ItemRepo:  itemModule,
		OrderRepo: orderModule,
	})
	return &serviceObj
}
