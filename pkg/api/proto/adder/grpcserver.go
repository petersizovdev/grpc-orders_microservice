package adder

import (
	"context"
	api "grpc-orders_microservice/pkg/api/proto"
)

type GRPCServer struct {
	api.UnimplementedAdderServer
}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}