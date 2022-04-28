package api

import "fmt"

func NewKV(key string, value interface{}) *KV {
	switch value.(type) {
	case string:
		return &KV{Key: key, Value: &KV_String_{String_: fmt.Sprint(value)}}
	case float64:
		v, ok := value.(float64)
		if !ok {
			return nil
		}
		return &KV{Key: key, Value: &KV_Float64{Float64: v}}
	case float32:
		v, ok := value.(float32)
		if !ok {
			return nil
		}
		return &KV{Key: key, Value: &KV_Float32{Float32: v}}
	case int, int32, int64, int8, int16:
		v, ok := value.(int64)
		if !ok {
			return nil
		}
		return &KV{Key: key, Value: &KV_Integer{Integer: v}}
	case bool:
		v, ok := value.(bool)
		if !ok {
			return nil
		}
		return &KV{Key: key, Value: &KV_Bool{Bool: v}}
	default:
		return &KV{Key: key, Value: &KV_String_{String_: fmt.Sprintf("%v", value)}}
	}
}
