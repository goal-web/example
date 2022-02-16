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
					"k":        .01,
					"filepath": "/Users/qbhy/project/go/goal-web/example/storage/default",
				},
				"users": contracts.Fields{
					"driver": "redis",
					"size":   100000,
					"k":      .01,
					"key":    "bloomfilter:{name}", // {name} 表示该 filter 的key，这里是 users
					//"connection": "cache", // redis 连接
				},
			},
		}
	}
}
