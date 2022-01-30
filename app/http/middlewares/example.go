package middlewares

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/http"
	"github.com/goal-web/pipeline"
)

func Example(request contracts.HttpRequest, next pipeline.Pipe) interface{} {
	fmt.Println("controller before")

	result := next(request)

	fmt.Println("controller after")
	return http.JsonResponse(contracts.Fields{
		"result": result,
	}, 200)
}
