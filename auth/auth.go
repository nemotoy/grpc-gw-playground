package auth

import (
	"context"
)

// Authenticator is an interface for assigning grpc-ecosystem/go-grpc-middleware/auth.
type Authenticator interface {
	Auth(ctx context.Context) (context.Context, error)
}
