package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
	"io/ioutil"
	"net/http"
	"strings"
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
* Date: 12.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Request struct {
	Endpoint    string
	httpRequest *http.Request
	httpClient  *http.Client
	Response    []byte
}

func New(Type networkAPI.Type, endpoint string) *Request {
	req := new(Request)
	req.httpRequest = new(http.Request)
	req.httpClient = new(http.Client)
	switch Type {
	case networkAPI.Type_Post:
		req.httpRequest.Method = http.MethodPost
	case networkAPI.Type_Get:
		req.httpRequest.Method = http.MethodGet
	}
	req.Endpoint = endpoint
	return req
}

func (req *Request) SetAuth(auth *networkAPI.Auth) error {
	switch auth.Type {
	case networkAPI.AuthType_None:
		return nil
	case networkAPI.AuthType_StaticToken:
	case networkAPI.AuthType_UsernamePassword:
	default:
		return errors.New("auth type not supported")
	}
	return nil
}

func (req *Request) SetHeaders(headers []*networkAPI.KV) error {
	for _, header := range headers {
		req.httpRequest.Header.Set(header.Key, header.Value)
	}
	return nil
}

func (req *Request) SetParams(params []*networkAPI.KV) error {
	switch req.httpRequest.Method {
	case http.MethodPost:
		if err := req.marshalParams(params); err != nil {
			return err
		}
	case http.MethodGet:
		req.appendParamsToEndpoint(params)
	}
	return nil
}

func (req *Request) Do() ([]byte, error) {
	resp, err := req.httpClient.Do(req.httpRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		return ioutil.ReadAll(resp.Body)
	default:
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}
}

func (req *Request) appendParamsToEndpoint(params []*networkAPI.KV) {
	req.Endpoint = strings.Trim(strings.TrimSuffix(strings.TrimSpace(req.Endpoint), "/"), "?")
	req.Endpoint = req.Endpoint + "?"
	for _, param := range params {
		if param.Key != "" && param.Value != "" {
			req.Endpoint = fmt.Sprintf("%s%s=%s&", req.Endpoint, param.Key, param.Value)
		}
	}
	req.Endpoint = strings.Trim(strings.TrimSuffix(strings.TrimSpace(req.Endpoint), "?"), "&")
}

func (req *Request) marshalParams(params []*networkAPI.KV) error {
	bodyMap := make(map[string]string)
	for _, param := range params {
		bodyMap[param.Key] = param.Value
	}
	jsonBody, err := json.Marshal(bodyMap)
	if err != nil {
		return err
	}
	req.httpRequest.Body = ioutil.NopCloser(bytes.NewReader(jsonBody))
	return nil
}
