package errors

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 14.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

func New(ctx context.Context, code codes.Code) *Error {
	return NewWithSlug(ctx, code, code.String())
}

func NewWithSlug(ctx context.Context, code codes.Code, slug string) *Error {
	return &Error{
		s:     status.Newf(code, slug),
		isRTL: false,
	}
}
