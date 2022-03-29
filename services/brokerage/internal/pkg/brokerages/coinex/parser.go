package coinex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	"go/types"
	"reflect"
)

type responseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func parseResponse(input string, response interface{}) error {
	tmp := new(responseModel)
	if err := json.Unmarshal([]byte(input), tmp); err != nil {
		return err
	}
	if tmp.Code != 0 {
		return errors.New(tmp.Message)
	}
	switch tmp.Data.(type) {
	case []interface{}:
		mapper.Slice(tmp.Data, response)
	case types.Struct:
	case map[string]interface{}:
		data, err := json.Marshal(tmp.Data)
		if err != nil {
			return err
		}
		return json.Unmarshal(data, response)
	default:
		return errors.New(fmt.Sprintf("cast data failed: %v", reflect.TypeOf(tmp.Data)))
	}
	return nil
}
