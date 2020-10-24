package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nemotoy/grpc-gw-playground/infra"
	pb "github.com/nemotoy/grpc-gw-playground/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", infra.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	md := metadata.New(map[string]string{})
	md["access_key"] = []string{"key1"}
	md["test_param"] = []string{"val1", "val2"}
	ctx = metadata.NewOutgoingContext(ctx, md)

	r, err := c.GetUser(ctx, &pb.UserRequest{Id: int64(1)})
	if err != nil {
		log.Printf("failed to get a user: %v", err)
	}

	log.Printf("Response: %v", r)
}
