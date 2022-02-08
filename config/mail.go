package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/email"
)

func init() {
	configs["mail"] = func(env contracts.Env) interface{} {
		return email.Config{
			Default: "default",
			Mailers: map[string]contracts.Fields{
				"default": {
					"host":     env.GetString("mail.host"),
					"port":     env.GetString("mail.port"),
					"username": env.GetString("mail.username"),
					"password": env.GetString("mail.password"),
				},
			},
		}
	}
}
