package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/jobs"
	"github.com/goal-web/serialization"
)

func init() {
	configs["serialization"] = func(env contracts.Env) interface{} {
		return serialization.Config{
			Default: "json",
			Class: []contracts.Class{
				jobs.DemoClass,
			},
		}
	}
}
