package api

import (
	"database/sql/driver"
	"encoding/json"
)

func (IndicatorType) InRange(v interface{}) bool {
	i, ok := IndicatorType_value[v.(IndicatorType).String()]
	return ok && i > 0
}
func (d *IndicatorType) Scan(value interface{}) error {
	*d = IndicatorType(IndicatorType_value[value.(string)])
	return nil
}
func (d IndicatorType) Value() (driver.Value, error) {
	return d.String(), nil
}
func (d *IndicatorType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*d = IndicatorType(IndicatorType_value[str])
	return nil
}
func (d IndicatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

//--------------------------------------------------------
func (Source) InRange(v interface{}) bool {
	i, ok := Source_value[v.(Source).String()]
	return ok && i > 0
}
func (d *Source) Scan(value interface{}) error {
	*d = Source(Source_value[value.(string)])
	return nil
}
func (d Source) Value() (driver.Value, error) {
	return d.String(), nil
}
func (d *Source) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*d = Source(Source_value[str])
	return nil
}
func (d Source) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
