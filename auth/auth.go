package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Auth struct {
	Keys map[string]string
}

func (a *Auth) Auth(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "error")
	}

	fmt.Println(md)

	keys := md.Get("access_key")
	if len(keys) < 1 {
		return nil, status.Error(codes.Unauthenticated, "key is not found")
	}

	key := keys[0]

	v, ok := a.Keys[key]
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "does not exist key")
	}

	fmt.Println(v)

	return ctx, nil
}
