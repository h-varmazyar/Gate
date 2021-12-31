package coinex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/pkg/mapper"
	"go/types"
	"reflect"
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
* Date: 26.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

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
		mapper.Slice(tmp.Data, response)
	case types.Struct:
		mapper.Struct(tmp.Data, response)
	case map[string]interface{}:
		response = tmp.Data.(map[string]interface{})
	default:
		return errors.New(fmt.Sprintf("cast data failed: %v", reflect.TypeOf(tmp.Data)))
	}
	return nil
}
