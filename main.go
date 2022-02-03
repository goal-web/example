package main

import (
	"github.com/goal-web/application"
	"github.com/goal-web/auth"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
	"github.com/goal-web/console"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	console2 "github.com/goal-web/example/app/console"
	"github.com/goal-web/example/app/exceptions"
	"github.com/goal-web/example/app/providers"
	config2 "github.com/goal-web/example/config"
	"github.com/goal-web/example/routes"
	"github.com/goal-web/filesystem"
	"github.com/goal-web/goal/signal"
	"github.com/goal-web/hashing"
	"github.com/goal-web/http"
	"github.com/goal-web/redis"
	"github.com/goal-web/session"
	"os"
)

func main() {
	app := application.Singleton()
	path, _ := os.Getwd()
	app.Instance("path", path)

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})

	app.RegisterServices(
		&config.ServiceProvider{
			Env:             os.Getenv("env"),
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
		&http.ServiceProvider{RouteCollectors: []interface{}{
			func(router contracts.Router) {
				router.Static("/", "public")
			},
			// 路由收集器
			routes.ApiRoutes,
		}},
		&console.ServiceProvider{
			ConsoleProvider: console2.NewKernel,
		},
		providers.AppServiceProvider{},
	)

	app.Call(func(console3 contracts.Console, input contracts.ConsoleInput) {
		console3.Run(input)
	})
}
