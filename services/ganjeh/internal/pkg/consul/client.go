package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
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
* Date: 17.11.21
* Github: https://github.com/h-varmazyar
* Email: hossein.varmazyar@yahoo.com
**/

var (
	client *api.Client
)

func init() {
	var err error
	client, err = api.NewClient(api.DefaultConfig())
	if err != nil {
		log.WithError(err).Panic("loading consul failed")
	}
}

func formatKey(inputs ...string) string {
	if len(inputs) == 0 {
		return ""
	}
	key := inputs[0]
	for _, input := range inputs[1:] {
		if input != "" {
			key = fmt.Sprintf("%s/%s", key, input)
		}
	}
	return key
}
