package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/filesystem"
	"os"
)

func init() {
	configs["filesystem"] = func(env contracts.Env) interface{} {
		return filesystem.Config{
			Default: env.GetString("filesystem"),
			Disks: map[string]contracts.Fields{
				"local": {
					"driver": "local",
					"name":   "local",
					"root":   env.GetString("filesystem.root"),
					"perm":   os.ModePerm,
				},
				"qiniu": {
					"driver":     "qiniu",
					"bucket":     env.GetString("qiniu.bucket"),
					"access_key": env.GetString("qiniu.access.key"),
					"secret_key": env.GetString("qiniu.secret.key"),
				},
			},
		}
	}
}
