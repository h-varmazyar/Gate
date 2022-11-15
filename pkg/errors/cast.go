package errors

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcStatus interface {
	GRPCStatus() *status.Status
}

func Cast(ctx context.Context, err error) *Error {
	switch e := err.(type) {
	case nil:
		return nil
	case *Error:
		return e
	case grpcStatus:
		return &Error{s: e.GRPCStatus()}
	}
	return New(ctx, codes.Unknown).
		AddDetails(err.Error())
}
