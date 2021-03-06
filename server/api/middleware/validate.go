package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) ValidatorInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)
	validator, ok := req.(interface{ Validate() error })
	if !ok {
		return handler(ctx, req)
	}
	if err := validator.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return handler(ctx, req)
}
