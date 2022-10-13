package api

import (
	"database/sql/driver"
	"encoding/json"
)

func (AuthType) InRange(v interface{}) bool {
	i, ok := AuthType_value[v.(AuthType).String()]
	return ok && i > 0
}
func (d *AuthType) Scan(value interface{}) error {
	*d = AuthType(AuthType_value[value.(string)])
	return nil
}
func (d AuthType) Value() (driver.Value, error) {
	return d.String(), nil
}
func (d *AuthType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*d = AuthType(AuthType_value[str])
	return nil
}
func (d AuthType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

//------------------------------------------------
func (Status) InRange(v interface{}) bool {
	i, ok := Status_value[v.(Status).String()]
	return ok && i > 0
}
func (d *Status) Scan(value interface{}) error {
	*d = Status(Status_value[value.(string)])
	return nil
}
func (d Status) Value() (driver.Value, error) {
	return d.String(), nil
}
func (d *Status) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*d = Status(Status_value[str])
	return nil
}
func (d Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
