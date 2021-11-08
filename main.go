package main

import (
	"fmt"
	"github.com/qbhy/goal/cache"
	"github.com/qbhy/goal/config"
	"github.com/qbhy/goal/container"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/encryption"
	"github.com/qbhy/goal/events"
	exceptions2 "github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/filesystemt"
	"github.com/qbhy/goal/hashing"
	"github.com/qbhy/goal/http/routes"
	"github.com/qbhy/goal/redis"
	"github.com/qbhy/goal/utils"
	"goal-example/app/exceptions"
	routes2 "goal-example/routes"
	"os"
)

func main() {

	// 初始化容器
	app := container.Singleton()

	// 获取当前目录
	path, _ := os.Getwd()

	// 设置运行目录
	app.Instance("path", path)

	// 注册异常处理器
	app.ProvideSingleton(func() contracts.ExceptionHandler {
		return exceptions.ExceptionHandler{
			DefaultExceptionHandler: exceptions2.NewDefaultHandler(exceptions.DontReportExceptions),
		}
	})

	// 注册各种服务
	app.RegisterServices(
		// 配置服务
		config.ServiceProvider{
			Paths: []string{path},
			Env:   os.Getenv("APP_ENV"),
		},
		// 加密 服务
		encryption.ServiceProvider{},
		// 异常处理服务，自定义了
		//exceptions.ServiceProvider{DontReportExceptions: []contracts.Exception{}},
		// http 路由服务，实例化路由组件
		routes.ServiceProvider{},
		// 事件调度服务
		events.ServiceProvider{},
		// 文件系统服务
		filesystemt.ServiceProvider{},
		// redis 服务
		redis.ServiceProvider{},
		// cache 服务
		cache.ServiceProvider{},
		// 哈希服务
		hashing.ServiceProvider{},
	)

	app.Call(routes2.RegisterV1Routes)


	app.Call(func(cache contracts.CacheFactory) {
		fmt.Println("cache get", cache.Store().Get("a"))
	})

	app.Call(func(router contracts.Router, conf contracts.Config) {
		fmt.Println(
			// 启动 http server，底层是 echo/v4 性能强悍
			router.Start(
				utils.StringOr(
					conf.GetString("server.address"),
					fmt.Sprintf("%s:%s",
						conf.GetString("server.host"),
						utils.StringOr(conf.GetString("server.port"), "8000"),
					),
				),
			),
		)

	})
}
