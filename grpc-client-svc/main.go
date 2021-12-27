package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	pb "mango.im/grpc-lp/api"
	"os"
	"time"
)

func main() {
	serverHost := os.Getenv("SERVER_HOST")
	if serverHost == "" {
		serverHost = ":8081"
	}
	conn, err := grpc.Dial(serverHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductServiceClient(conn)

	for i := 0; i < 1000; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		guid, _ := uuid.NewUUID()
		r, err := c.GetProduct(ctx, &pb.GetProductRequest{ProductId: guid.String()})
		if err != nil {
			log.Fatalf("could not get product: %v", err)
		}
		log.Printf("product name: %s", r.GetName())
	}
}
