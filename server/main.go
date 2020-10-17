package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nemotoy/grpc-gw-playground/proto/user"
	"google.golang.org/grpc"
)

const (
	port = 50051
)

type userServer struct{}

func newUserServer() user.UserServiceServer {
	return &userServer{}
}

// GetUser implements userServiceSever.
func (s *userServer) GetUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	// do something
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	defer s.GracefulStop()
	user.RegisterUserServiceServer(s, newUserServer())

	log.Println("server starts: ", port)

	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
