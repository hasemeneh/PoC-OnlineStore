package service

import (
	"log"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-OnlineStore/helper/redis"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/definitions"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/domain/cart"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/domain/item"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/interface/order"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/interface/warehouse"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/usecase/checkout"
)

type Service struct {
	cfg     *models.MainConfig
	Usecase struct {
		Checkout definitions.CheckoutDefinition
	}
}

func New(cfg *models.MainConfig) *Service {
	db := database.New().Connect(cfg.DBConnectionString)
	redisInterface := redis.New(cfg.RedisConnection, "")

	cartRepo := cart.New(db, redisInterface)
	itemRepo := item.New(db)
	warehouseInterface := warehouse.New(&warehouse.Option{
		Address: cfg.WarehouseAddress,
		Fatalln: log.Fatalln,
	})
	orderInterface := order.New(&order.Option{
		Address: cfg.OrderAddress,
		Fatalln: log.Fatalln,
	})
	serviceObj := Service{
		cfg: cfg,
	}
	serviceObj.Usecase.Checkout = checkout.New(&checkout.Option{
		ItemRepo:  itemRepo,
		CartRepo:  cartRepo,
		Order:     orderInterface,
		Warehouse: warehouseInterface,
	})

	return &serviceObj
}
