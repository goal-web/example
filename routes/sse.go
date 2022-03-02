package routes

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/http/sse"
)

func Sse(router contracts.Router) {
	router.Get("/sse", sse.Default())

	router.Get("/send-sse", func(sse contracts.Sse, request contracts.HttpRequest) error {
		return sse.Send(uint64(request.GetInt64("fd")), request.GetString("msg"))
	})

	router.Get("/close-sse", func(sse contracts.Sse, request contracts.HttpRequest) error {
		return sse.Close(uint64(request.GetInt64("fd")))
	})
}
