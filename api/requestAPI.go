package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"
)

type RequestType int

type Request struct {
	Type     RequestType
	Endpoint string
	Headers  map[string][]string
	Params   map[string]interface{}
}

type Response struct {
	Code    int
	Status  string
	Body    []byte
	Headers map[string][]string
}

const (
	GET RequestType = iota
	PUT
	POST
	PATCH
	DELETE
)

func (r *Request) Execute() (*Response, error) {
	switch r.Type {
	case GET:
		return r.get()
	case PUT:
		return nil, errors.New("method not supported yet")
	case POST:
		return r.post()
	case PATCH:
		return nil, errors.New("method not supported yet")
	case DELETE:
		return nil, errors.New("method not supported yet")
	}
	return nil, errors.New("please declare valid method")
}

func (r *Request) get() (*Response, error) {
	params := "?"
	for key, value := range r.Params {
		params += fmt.Sprintf("%s=%v", key, value)
	}
	endpoint := strings.TrimSuffix(r.Endpoint, "/")
	endpoint += params
	endpoint = strings.TrimSuffix(r.Endpoint, "?")

	url, err := url2.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: r.Headers,
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Code:    resp.StatusCode,
		Status:  resp.Status,
		Body:    data,
		Headers: resp.Header,
	}, nil
}

func (r *Request) post() (*Response, error) {
	url, err := url2.Parse(r.Endpoint)
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(r.Params)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method: http.MethodPost,
		URL:    url,
		Header: r.Headers,
		Body:   ioutil.NopCloser(bytes.NewReader(body)),
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Code:    resp.StatusCode,
		Status:  resp.Status,
		Body:    data,
		Headers: resp.Header,
	}, nil
}
