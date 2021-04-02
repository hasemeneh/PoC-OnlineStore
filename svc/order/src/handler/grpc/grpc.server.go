package grpc

import (
	"log"
	"net"
	"os"
	"syscall"

	pb "github.com/hasemeneh/PoC-OnlineStore/svc/grpc/order/protos"

	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	Service  *service.Service
	server   *grpc.Server
	GrpcPort string
	Fatalln  logger
}
type Option struct {
	GrpcPort string
	Service  *service.Service
}
type GrpcServer interface {
	Serve()
	CatchSignal(s os.Signal)
}
type logger func(v ...interface{})

func (se *server) Serve() {
	lis, err := net.Listen("tcp", se.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	se.server = s
	pb.RegisterOrderServer(s, se)
	reflection.Register(s)
	log.Println("Serving warehouse GRPC on ", se.GrpcPort)
	if err := s.Serve(lis); err != nil {
		se.Fatalln("failed to serve: ", err)
	}

}
func (se *server) CatchSignal(s os.Signal) {
	log.Println("grpc service got signal:", s)
	if s == syscall.SIGHUP || s == syscall.SIGTERM {
		se.server.GracefulStop()
	}
}
func New(o *Option) GrpcServer {

	return &server{
		Service:  o.Service,
		GrpcPort: o.GrpcPort,
		Fatalln:  log.Fatalln,
	}
}
