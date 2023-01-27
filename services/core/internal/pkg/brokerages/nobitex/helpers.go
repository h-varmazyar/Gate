package nobitex

import (
	"fmt"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
)

func (r *Requests) generateAuthorization() string {
	return fmt.Sprintf("Token %v", r.Auth.NobitexToken)
}

func parseValue(param *networkAPI.KV) string {
	value := ""
	switch v := param.Value.(type) {
	case *networkAPI.KV_String_:
		value = v.String_
	case *networkAPI.KV_Bool:
		value = fmt.Sprint(v.Bool)
	case *networkAPI.KV_Float32:
		value = fmt.Sprint(v.Float32)
	case *networkAPI.KV_Float64:
		value = fmt.Sprint(v.Float64)
	case *networkAPI.KV_Integer:
		value = fmt.Sprint(v.Integer)
	}
	return value
}
