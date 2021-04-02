package order

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
	order_GRPC "github.com/hasemeneh/PoC-OnlineStore/svc/grpc/order/protos"

	"google.golang.org/grpc"
)

type OrderInterface struct {
	client order_GRPC.OrderClient
}
type OrderManager interface {
	Checkout(ctx context.Context, in *models.OrderCreationRequest) (*models.OrderCreationResponse, error)
}
type Option struct {
	Address string
	Fatalln logger
}
type logger func(v ...interface{})

func New(o *Option) *OrderInterface {
	conn, err := grpc.Dial(o.Address, grpc.WithInsecure())
	if err != nil {
		o.Fatalln("did not connect: ", err)
	}
	client := order_GRPC.NewOrderClient(conn)
	return &OrderInterface{
		client: client,
	}
}
