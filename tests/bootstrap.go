package tests

import (
	"github.com/goal-web/application"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	config2 "github.com/goal-web/example/config"
	"github.com/goal-web/filesystem"
	"github.com/goal-web/goal/auth"
	"github.com/goal-web/goal/exceptions"
	"github.com/goal-web/goal/session"
	"github.com/goal-web/goal/signal"
	"github.com/goal-web/hashing"
	"github.com/goal-web/redis"
	"github.com/goal-web/supports/logs"
)

func initApp(path string) contracts.Application {
	env := "testing"
	app := application.Singleton()
	app.Instance("path", path)

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.DefaultExceptionHandler{}
	})

	app.RegisterServices(
		&config.ServiceProvider{
			Env:             env,
			Paths:           []string{path},
			Sep:             "=",
			ConfigProviders: config2.Configs(),
		},
		hashing.ServiceProvider{},
		encryption.ServiceProvider{},
		filesystem.ServiceProvider{},
		events.ServiceProvider{},
		redis.ServiceProvider{},
		cache.ServiceProvider{},
		&signal.ServiceProvider{},
		&session.ServiceProvider{},
		auth.ServiceProvider{},
		&database.ServiceProvider{},
		//&http.ServiceProvider{RouteCollectors: []interface{}{
		//	// 路由收集器
		//	routes.V1Routes,
		//}},
	)

	go func() {
		if errors := app.Start(); len(errors) > 0 {
			logs.WithField("errors", errors).Fatal("goal 启动异常!")
		}
	}()
	return app
}
