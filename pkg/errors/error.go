package errors

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
	"reflect"
	"strings"
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

type Error struct {
	s     *status.Status
	isRTL bool
}

func (e Error) Error() string {
	return fmt.Sprintf("%s(%d):%s", e.s.Message(), e.s.Code(), strings.Join(e.Details(), "\n"))
}

func (e Error) Details() []string {
	details := make([]string, 0)
	if reflect.TypeOf(e.s.Details()) == reflect.TypeOf(reflect.Slice) {
		for _, detail := range e.s.Details() {
			e, ok := detail.(*errdetails.ErrorInfo)
			if ok {
				details = append(details, e.Reason)
			}
		}
	}
	return details
}

func (e Error) AddDetails(details ...string) *Error {
	var errDetails []proto.Message
	for _, detail := range details {
		errDetails = append(errDetails, &errdetails.ErrorInfo{
			Reason: detail,
		})
	}
	s, _ := e.s.WithDetails(errDetails...)
	return &Error{s: s}
}
