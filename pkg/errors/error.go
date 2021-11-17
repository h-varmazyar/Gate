package errors

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type Error struct {
	s *status.Status
}

func (e Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&Model{
		Message: e.Message(),
		Details: e.Details(),
		Code:    uint32(e.Code()),
	})
}
func (e Error) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	return encoder.EncodeElement(&Model{
		Message: e.Message(),
		Details: e.Details(),
		Code:    uint32(e.Code()),
	}, start)
}
func (e Error) MarshalYAML() (interface{}, error) {
	return &Model{
		Message: e.Message(),
		Details: e.Details(),
		Code:    uint32(e.Code()),
	}, nil
}

func (e Error) Message() string {
	return e.s.Message()
}

func (e Error) Details() []string {
	var details []string
	for _, detail := range e.s.Details() {
		ei, ok := detail.(*errdetails.ErrorInfo)
		if !ok {
			continue
		}
		details = append(details, ei.Reason)
	}
	return details
}

func (e Error) Error() string {
	return fmt.Sprintf("%s | %s | %s",
		strings.ToUpper(e.s.Code().String()), e.s.Message(), strings.Join(e.Details(), ", "))
}

func (e Error) GRPCStatus() *status.Status {
	return e.s
}

func (e Error) Code() codes.Code {
	return e.s.Code()
}

func (e Error) HttpStatus() int {
	return runtime.HTTPStatusFromCode(e.s.Code())
}

func (e Error) IsHandled() bool {
	return IsHandledCode(e.s.Code())
}

func (e Error) AddDetailF(format string, args ...interface{}) *Error {
	return e.AddDetails(fmt.Sprintf(format, args...))
}

func (e Error) AddDetails(details ...string) *Error {
	var errDetails []proto.Message
	for _, detail := range details {
		errDetails = append(errDetails, &errdetails.ErrorInfo{
			Reason: detail,
		})
	}
	s, _ := e.s.WithDetails(errDetails...)
	return &Error{s}
}
