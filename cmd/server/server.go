package server

import (
	"grpc-orders_microservice/pkg/api/proto"
	"grpc-orders_microservice/pkg/api/proto/adder"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}