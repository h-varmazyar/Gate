package consul

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	ganjehAPI "github.com/mrNobody95/Gate/services/ganjeh/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
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
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

func Set(variable *ganjehAPI.Variable) error {
	val, err := json.Marshal(variable)
	if err != nil {
		return err
	}
	pair := &api.KVPair{
		Key:   formatKey(variable.Namespace, variable.Key),
		Value: val,
	}
	_, err = client.KV().Put(pair, nil)
	return err
}

func Get(ctx context.Context, namespace, key string) (*ganjehAPI.Variable, error) {
	pair, _, err := client.KV().Get(formatKey(namespace, key), nil)
	if err != nil {
		return nil, err
	}
	if pair == nil {
		return nil, errors.New(ctx, codes.NotFound)
	}
	variable := new(ganjehAPI.Variable)
	return variable, json.Unmarshal(pair.Value, variable)
}

func GetList(ctx context.Context, namespace string) (*ganjehAPI.Variables, error) {
	kto := api.KVTxnOps{}
	kto = append(kto, &api.KVTxnOp{
		Verb: api.KVGetTree,
		Key:  formatKey(namespace),
	})
	_, resp, _, err := client.KV().Txn(kto, nil)
	if err != nil {
		return nil, err
	}
	if resp == nil || len(resp.Results) == 0 {
		return nil, errors.New(ctx, codes.NotFound)
	}
	variables := new(ganjehAPI.Variables)
	for _, result := range resp.Results {
		variable := new(ganjehAPI.Variable)
		if err := json.Unmarshal(result.Value, variable); err != nil {
			log.WithError(err).Errorf("parsing variable failed for: %s", string(result.Value))
			continue
		}
		variables.Variables = append(variables.Variables, variable)
	}
	return variables, nil
}
