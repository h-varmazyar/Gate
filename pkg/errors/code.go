package errors

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
)

func IsHandledCode(code codes.Code) bool {
	if status := runtime.HTTPStatusFromCode(code); status >= 500 && status <= 599 {
		return false
	}
	return true
}

func CodeToSlug(code codes.Code) string {
	return code.String()
}

func Code(ctx context.Context, err error) codes.Code {
	if e := Cast(ctx, err); e != nil {
		return e.Code()
	}
	return codes.OK
}
