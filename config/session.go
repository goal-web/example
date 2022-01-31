package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/session"
	"time"
)

func init() {
	configs["session"] = func(env contracts.Env) interface{} {
		return session.Config{
			Driver:     "cookie",
			Encrypt:    true,
			Domain:     env.GetString("session.domain"),
			Lifetime:   time.Duration(env.GetInt("session.lifetime")),
			Connection: "", // database、redis 用到
			Table:      "", // database 用到
			Name:       env.StringOption("session.domain", "goal_"),
		}
	}
}
