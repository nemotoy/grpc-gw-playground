package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	stub = map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}
)

type Authenticator interface {
	Auth(ctx context.Context) (context.Context, error)
}

type Auth struct {
	Keys map[string]string
}

func New() Authenticator {
	return &Auth{Keys: stub}
}

func (a *Auth) Auth(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "error")
	}

	fmt.Println("metadata: ", md)

	keys := md.Get("access_key")
	if len(keys) < 1 {
		return nil, status.Error(codes.Unauthenticated, "key is not found")
	}

	fmt.Println("keys: ", keys)

	key := keys[0]

	if _, ok := a.Keys[key]; !ok {
		return nil, status.Error(codes.Unauthenticated, "does not exist key")
	}

	return ctx, nil
}
