package rpc

import (
	"hexagonal-go/internal/ports"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", err)
	}
	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	// TODO: implementations
}
