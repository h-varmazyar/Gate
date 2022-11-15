package errors

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(ctx context.Context, code codes.Code) *Error {
	return &Error{status.New(code, CodeToSlug(code))}
}

func NewWithSlug(ctx context.Context, code codes.Code, slug string) *Error {
	return &Error{status.New(code, slug)}
}
