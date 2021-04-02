package warehouse

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
	warehouse_GRPC "github.com/hasemeneh/PoC-OnlineStore/svc/grpc/warehouse/protos"
	"google.golang.org/grpc"
)

type warehouseInterface struct {
	client warehouse_GRPC.WarehouseClient
}
type WarehouseManager interface {
	ValidateStock(ctx context.Context, req *models.ProductRequest) (*models.ProductResponse, error)
}
type Option struct {
	Address string
	Fatalln logger
}
type logger func(v ...interface{})

func New(o *Option) *warehouseInterface {
	conn, err := grpc.Dial(o.Address, grpc.WithInsecure())
	if err != nil {
		o.Fatalln("did not connect: ", err)
	}
	client := warehouse_GRPC.NewWarehouseClient(conn)

	return &warehouseInterface{
		client: client,
	}
}
