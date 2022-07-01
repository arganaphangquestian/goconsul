package utils

import (
	"log"
	"strconv"

	"github.com/hashicorp/consul/api"
)

type FeatureFlag struct {
	Client *api.Client
}

func (f FeatureFlag) Get(path string) bool {
	pair, _, err := f.Client.KV().Get(path, nil)
	if err != nil {
		log.Println(err)
		return false
	}
	p, err := strconv.ParseBool(string(pair.Value))
	if err != nil {
		return false
	}
	return p
}
