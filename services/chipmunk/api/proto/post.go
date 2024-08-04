package proto

import (
	"database/sql/driver"
	"encoding/json"
)

func (Provider) InRange(v interface{}) bool {
	i, ok := Provider_value[v.(Provider).String()]
	return ok && i > 0
}
func (d *Provider) Scan(value interface{}) error {
	*d = Provider(Provider_value[value.(string)])
	return nil
}
func (d Provider) Value() (driver.Value, error) {
	return d.String(), nil
}
func (d *Provider) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*d = Provider(Provider_value[str])
	return nil
}
func (d Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (Polarity) InRange(v interface{}) bool {
	i, ok := Polarity_value[v.(Provider).String()]
	return ok && i > 0
}
func (d *Polarity) Scan(value interface{}) error {
	*d = Polarity(Polarity_value[value.(string)])
	return nil
}
func (d Polarity) Value() (driver.Value, error) {
	return d.String(), nil
}
func (d *Polarity) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*d = Polarity(Polarity_value[str])
	return nil
}
func (d Polarity) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
