package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	pb "mango.im/grpc-lp/api"
	"math/rand"
	"net"
	"time"
)

type server struct {
	pb.UnsafeProductServiceServer
}

func (s *server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	log.Printf("product id: %v", req.GetProductId())
	return &pb.GetProductResponse{
		Name:        "name_" + req.GetProductId(),
		Description: "description_" + req.ProductId,
		Price:       rand.Float64()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//s := grpc.NewServer()
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)

	pb.RegisterProductServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
