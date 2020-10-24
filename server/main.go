package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/nemotoy/grpc-gw-playground/auth"
	"github.com/nemotoy/grpc-gw-playground/infra"
	pb "github.com/nemotoy/grpc-gw-playground/proto/user"
	"google.golang.org/grpc"
)

var (
	keysStub = map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}
)

type userServer struct{}

func newUserServer() pb.UserServiceServer {
	return &userServer{}
}

// GetUser implements userServiceSever.
func (s *userServer) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	// do something
	return new(pb.UserResponse), nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", infra.Port))
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

	pb.RegisterUserServiceServer(s, newUserServer())

	log.Printf("server starts: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
