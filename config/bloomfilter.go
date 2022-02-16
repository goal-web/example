package config

import (
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/contracts"
)

func init() {
	configs["bloomfilter"] = func(env contracts.Env) interface{} {
		return bloomfilter.Config{
			Default: "default",
			Filters: bloomfilter.Filters{
				"default": contracts.Fields{
					"driver":   "file",
					"size":     100000,
					"k":        0,
					"filepath": "/Users/qbhy/project/go/goal-web/example/storage/default",
				},
			},
		}
	}
}
