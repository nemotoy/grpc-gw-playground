package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/nemotoy/grpc-gw-playground/auth"
	"github.com/nemotoy/grpc-gw-playground/proto/user"
	"google.golang.org/grpc"
)

const (
	port = 50051
)

var (
	keysStub = map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}
)

type userServer struct{}

func newUserServer() user.UserServiceServer {
	return &userServer{}
}

// GetUser implements userServiceSever.
func (s *userServer) GetUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	// do something
	return new(user.UserResponse), nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authImpl := &auth.Auth{Keys: keysStub}
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_auth.UnaryServerInterceptor(authImpl.Auth),
		),
	)

	defer s.GracefulStop()

	user.RegisterUserServiceServer(s, newUserServer())

	log.Println("server starts: ", port)

	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
