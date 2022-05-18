package api

import (
	"database/sql/driver"
	"encoding/json"
)

func (Platform) InRange(v interface{}) bool {
	_, ok := Platform_value[v.(Platform).String()]
	return ok
}
func (x *Platform) Scan(value interface{}) error {
	*x = Platform(Platform_value[value.(string)])
	return nil
}
func (x Platform) Value() (driver.Value, error) {
	return x.String(), nil
}
func (x *Platform) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*x = Platform(Platform_value[str])
	return nil
}
func (x Platform) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
