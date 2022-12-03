package proto

import (
	"database/sql/driver"
	"encoding/json"
)

func (RateLimiterType) InRange(v interface{}) bool {
	i, ok := RateLimiterType_value[v.(RateLimiterType).String()]
	return ok && i > 0
}
func (d *RateLimiterType) Scan(value interface{}) error {
	*d = RateLimiterType(RateLimiterType_value[value.(string)])
	return nil
}
func (d RateLimiterType) Value() (driver.Value, error) {
	return d.String(), nil
}
func (d *RateLimiterType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*d = RateLimiterType(RateLimiterType_value[str])
	return nil
}
func (d RateLimiterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
