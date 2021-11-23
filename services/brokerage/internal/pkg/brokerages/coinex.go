package brokerages

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
	"io"
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
* Date: 19.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Coinex struct {
	Auth *api.Auth
}

func (c *Coinex) WalletList(runner func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error)) ([]*repository.Wallet, error) {
	request := new(networkAPI.Request)
	request.Type = networkAPI.Type_GET
	request.Endpoint = "https://api.coinex.com/v1/balance/info"
	fmt.Println("acc:", c.Auth.AccessID)
	request.Params = []*networkAPI.KV{
		{Key: "access_id", Value: c.Auth.AccessID},
		{Key: "tonce", Value: fmt.Sprintf("%d", time.Now().UnixNano()/1e6)},
	}
	request.Headers = []*networkAPI.KV{
		{Key: "authorization", Value: c.generateAuthorization(request.Params)},
		{Key: "tonce", Value: fmt.Sprintf("%d", time.Now().UnixNano()/1e6)},
	}
	resp, err := runner(context.Background(), request)
	if err != nil {
		return nil, err
	}
	fmt.Println("resp:", resp)
	return nil, nil
}

func (c *Coinex) generateAuthorization(params []*networkAPI.KV) string {
	urlParameters := url.Values{}
	for _, param := range params {
		urlParameters.Add(param.Key, param.Value)
	}
	queryParamsString := urlParameters.Encode()
	toEncodearamsString := queryParamsString + "&secrect=" + c.Auth.SecretKey
	w := md5.New()
	io.WriteString(w, toEncodearamsString)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))
	md5Str = strings.ToUpper(md5Str)
	return md5Str
}
