package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
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
		req.headers.Add(header.Key, header.Value)
	}
}

func (req *Request) AddQueryParams(params []*networkAPI.KV) {
	qParams := url.Values{}
	for _, param := range params {
		qParams.Add(param.Key, param.Value)
	}
	req.queryParams = qParams.Encode()
}

func (req *Request) AddBody(bodyParams []*networkAPI.KV) error {
	bodyMap := make(map[string]string)
	for _, param := range bodyParams {
		bodyMap[param.Key] = param.Value
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
