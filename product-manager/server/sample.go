package server

import (
	"context"
	"fmt"
	hellopb "yorushika-store/product-manager/pkg/grpc"
	"yorushika-store/product-manager/repository"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
	db *repository.Database
}

func NewMyServer(db *repository.Database) *myServer {
	return &myServer{db: db}
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	products, err := s.db.ListProducts()
	if err != nil {
		return nil, err
	}

	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s", products[0].Name),
	}, nil
}
