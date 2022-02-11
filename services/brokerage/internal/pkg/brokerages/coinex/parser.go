package coinex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/pkg/mapper"
	"go/types"
	"reflect"
)

type responseModel struct {
	Code    int
	Message string
	Data    interface{}
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
		mapper.Slice(tmp.Data, &response)
	case types.Struct:
		mapper.Struct(tmp.Data, response)
	case map[string]interface{}:
		response = tmp.Data.(map[string]interface{})
	default:
		return errors.New(fmt.Sprintf("cast data failed: %v", reflect.TypeOf(tmp.Data)))
	}
	return nil
}
