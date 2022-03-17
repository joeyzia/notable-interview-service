package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// AuthMiddleware returns a Basic Authentication middleware for a particular user and password.
func SomeMiddleware(anything string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// Do something fun
			return next(ctx, request)
		}
	}
}