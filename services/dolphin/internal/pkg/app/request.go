package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type KV struct {
	Key   string
	Value string
}

type Request struct {
	Endpoint   string
	Method     string
	Body       io.Reader
	Header     http.Header
	Param      url.Values
	httpClient *http.Client
	Response   []byte
}

func NewRequest(method string, endpoint ...string) *Request {

	return &Request{
		httpClient: new(http.Client),
		Method:     method,
		Endpoint:   strings.Join(endpoint, ""),
		Header:     http.Header{},
		Param:      url.Values{},
		Body:       new(bytes.Reader),
	}
}

func (req *Request) SetHeader(key, value string) {
	req.Header.Set(key, value)
}

func (req *Request) SetHeaders(headers []KV) error {
	for _, header := range headers {
		req.Header.Set(header.Key, header.Value)
	}
	return nil
}

func (req *Request) SetParams(params map[string]string) error {
	for _, value := range params {
		req.Endpoint = fmt.Sprint(req.Endpoint, "/", value)
	}
	return nil
}

func (req *Request) SetQueryParams(params map[string]string) error {
	for key, value := range params {
		req.Param.Add(key, value)
	}
	return nil
}

func (req *Request) SetJsonBody(body interface{}) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req.Body = bytes.NewReader(jsonBody)
	return nil
}

func (req *Request) SetBody(body io.Reader) error {
	req.Body = body
	return nil
}

func (req *Request) Do() ([]byte, error) {
	request, err := http.NewRequest(
		req.Method,
		req.Endpoint,
		req.Body,
	)
	if err != nil {
		return nil, err
	}
	URL, err := url.Parse(req.Endpoint)
	if err != nil {
		return nil, err
	}

	request.URL = URL
	request.URL.RawQuery = req.Param.Encode()
	request.Header = req.Header
	resp, err := req.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		return ioutil.ReadAll(resp.Body)
	default:
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status: %s(%s)", resp.Status, string(body))
	}
}
