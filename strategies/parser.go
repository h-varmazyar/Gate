package strategies

import (
	"encoding/json"
	"encoding/xml"
)

func ParseJson(input string) (*Strategy, error) {
	s := Strategy{}
	err := json.Unmarshal([]byte(input), &s)
	if err != nil {
		return nil, err
	}
	return &s, err
}

func ParseXml(input string) (*Strategy, error) {
	s := Strategy{}
	err := xml.Unmarshal([]byte(input), &s)
	if err != nil {
		return nil, err
	}
	return &s, err
}
