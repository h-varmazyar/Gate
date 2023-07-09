package requests

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	ContentTypeJson = "application/json"
	UserAgent       = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"
)

type Request struct {
	Endpoint    string
	httpClient  *http.Client
	headers     http.Header
	queryParams string
	body        *bytes.Buffer
	method      networkAPI.RequestMethod
	metadata    string
}

func New(input *networkAPI.Request, proxyURL *url.URL) (*Request, error) {
	requestTransport := new(http.Transport)

	if proxyURL != nil {
		requestTransport.Proxy = http.ProxyURL(proxyURL)
		requestTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	request := &Request{
		Endpoint: input.Endpoint,
		httpClient: &http.Client{
			Transport: requestTransport,
		},
		method:   input.Method,
		body:     new(bytes.Buffer),
		headers:  http.Header{},
		metadata: input.Metadata,
	}

	request.AddHeaders(input.Headers)
	switch input.Method {
	case networkAPI.Request_POST:
		err := request.SetBody(input.Params)
		if err != nil {
			return nil, err
		}
	case networkAPI.Request_GET:
		request.SetQueryParams(input.Params)
	}

	return request, nil
}

func (req *Request) AddHeaders(headers []*networkAPI.KV) {
	for _, header := range headers {
		req.headers.Add(header.Key, parseValue(header))
	}
}

func (req *Request) SetQueryParams(params []*networkAPI.KV) {
	qParams := url.Values{}
	for _, param := range params {
		qParams.Add(param.Key, parseValue(param))
	}
	req.queryParams = qParams.Encode()
}

func (req *Request) SetBody(bodyParams []*networkAPI.KV) error {
	bodyMap := make(map[string]string)
	for _, param := range bodyParams {
		bodyMap[param.Key] = parseValue(param)
	}
	jsonBody, err := json.Marshal(bodyMap)
	if err != nil {
		return err
	}
	req.body = bytes.NewBuffer(jsonBody)
	return nil
}

func (req *Request) Do() (*networkAPI.Response, error) {
	request, err := http.NewRequest(strings.ToUpper(req.method.String()), req.Endpoint, req.body)
	if err != nil {
		log.WithError(err).Errorf("failed to create request")
		return nil, err
	}
	request.Header = req.headers
	request.Header.Set("Content-Type", ContentTypeJson)
	request.Header.Set("User-Agent", UserAgent)
	if req.method == networkAPI.Request_GET {
		request.URL.RawQuery = req.queryParams
	}

	log.Infof(request.URL.String())
	response, err := req.httpClient.Do(request)
	if err != nil {
		log.WithError(err).Errorf("failed to make request")
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &networkAPI.Response{
		Code:     int32(response.StatusCode),
		Body:     string(body),
		Metadata: req.metadata,
		Method:   req.method,
	}, nil
}

func parseValue(param *networkAPI.KV) string {
	value := ""
	switch v := param.Value.(type) {
	case *networkAPI.KV_String_:
		value = v.String_
	case *networkAPI.KV_Bool:
		value = fmt.Sprint(v.Bool)
	case *networkAPI.KV_Float32:
		value = fmt.Sprint(v.Float32)
	case *networkAPI.KV_Float64:
		value = fmt.Sprint(v.Float64)
	case *networkAPI.KV_Integer:
		value = fmt.Sprint(v.Integer)
	}
	return value
}
