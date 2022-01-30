package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
)

func init() {
	configs["database"] = func(env contracts.Env) interface{} {
		return database.Config{
			Default: env.StringOption("db.connection", "mysql"),
			Connections: map[string]contracts.Fields{
				"sqlite": {
					"driver":   "sqlite",
					"database": env.GetString("sqlite.database"),
				},
				"mysql": {
					"driver":          "mysql",
					"host":            env.GetString("db.host"),
					"port":            env.GetString("db.port"),
					"database":        env.GetString("db.database"),
					"username":        env.GetString("db.username"),
					"password":        env.GetString("db.password"),
					"unix_socket":     env.GetString("db.unix_socket"),
					"charset":         env.StringOption("db.charset", "utf8mb4"),
					"collation":       env.StringOption("db.collation", "utf8mb4_unicode_ci"),
					"prefix":          env.GetString("db.prefix"),
					"strict":          env.GetBool("db.struct"),
					"max_connections": env.GetInt("db.max_connections"),
					"max_idles":       env.GetInt("db.max_idles"),
				},
				"pgsql": {
					"driver":          "postgres",
					"host":            env.GetString("db.pgsql.host"),
					"port":            env.GetString("db.pgsql.port"),
					"database":        env.GetString("db.pgsql.database"),
					"username":        env.GetString("db.pgsql.username"),
					"password":        env.GetString("db.pgsql.password"),
					"charset":         env.StringOption("db.pgsql.charset", "utf8mb4"),
					"prefix":          env.GetString("db.pgsql.prefix"),
					"schema":          env.StringOption("db.pgsql.schema", "public"),
					"sslmode":         env.StringOption("db.pgsql.sslmode", "disable"),
					"max_connections": env.GetInt("db.pgsql.max_connections"),
					"max_idles":       env.GetInt("db.pgsql.max_idles"),
				},
			},
		}
	}
}
