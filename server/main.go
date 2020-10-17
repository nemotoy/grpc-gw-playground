package main

import (
	"context"
	"log"
	"net"

	"github.com/nemotoy/grpc-gw-playground/proto/user"
	"google.golang.org/grpc"
)

type userServer struct{}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	defer s.GracefulStop()
	user.RegisterUserServiceServer(s, newUserServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func newUserServer() user.UserServiceServer {
	return &userServer{}
}

func (s *userServer) GetUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	// do something
	return nil, nil
}
