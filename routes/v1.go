package routes

import (
	"github.com/qbhy/goal/contracts"
	v1 "goal-example/app/controllers"
)

func RegisterV1Routes(router contracts.Router) {
	router.Get("/", v1.Hello)
	router.Get("/hashing", v1.Hashing)
	router.Get("/redis", v1.Redis)
	router.Get("/file", v1.Filesystem)
	router.Get("/cache", v1.Cache)
	router.Get("/di", v1.DiRequest)
}
