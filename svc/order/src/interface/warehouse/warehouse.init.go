package warehouse

import (
	"context"

	warehouse_GRPC "github.com/hasemeneh/PoC-OnlineStore/svc/grpc/warehouse/protos"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
	"google.golang.org/grpc"
)

type warehouseInterface struct {
	client warehouse_GRPC.WarehouseClient
}
type WarehouseManager interface {
	ClaimProduct(ctx context.Context, req *models.ClaimRequest) (*models.ClaimResponse, error)
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
