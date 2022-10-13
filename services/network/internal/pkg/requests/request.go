package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	method      networkAPI.Type
}

func New(Type networkAPI.Type, endpoint string) *Request {
	return &Request{
		Endpoint: endpoint,
		httpClient: &http.Client{
			Timeout: 20 * time.Second,
			Transport: &http.Transport{
				TLSHandshakeTimeout: 30 * time.Second,
			},
		},
		method:  Type,
		body:    new(bytes.Buffer),
		headers: http.Header{},
	}
}

func (req *Request) AddHeaders(headers []*networkAPI.KV) {
	for _, header := range headers {
		req.headers.Add(header.Key, fmt.Sprint(header.Value))
	}
}

func (req *Request) AddQueryParams(params []*networkAPI.KV) {
	qParams := url.Values{}
	log.Infof("params: %v", params)
	for _, param := range params {
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
		default:
			continue
		}
		qParams.Add(param.Key, value)
	}
	req.queryParams = qParams.Encode()

	log.Infof("params formated: %v", req.queryParams)
}

func (req *Request) AddBody(bodyParams []*networkAPI.KV) error {
	bodyMap := make(map[string]string)
	for _, param := range bodyParams {
		bodyMap[param.Key] = fmt.Sprint(param.Value)
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
		return nil, err
	}
	request.Header = req.headers
	request.Header.Set("Content-Type", ContentTypeJson)
	request.Header.Set("User-Agent", UserAgent)
	if req.method == networkAPI.Type_GET {
		request.URL.RawQuery = req.queryParams
	}
	response, err := req.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	fmt.Println(response.Request.URL)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &networkAPI.Response{
		Code:     int32(response.StatusCode),
		Response: string(body),
	}, nil
}
