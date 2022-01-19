package grpc

import (
	"log"
	"net"
	"github.tesla.com/chrzhang/go-hex-arch/internal/adapters/framework/left/grpc/pb"
	"github.tesla.com/chrzhang/go-hex-arch/internal/ports"
	"google.golang.org/grpc"
)

//adapter for grpc port
type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

//implement run method from port - which starts GRPC service
func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp:", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil{
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}