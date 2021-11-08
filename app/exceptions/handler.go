package exceptions

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/http"
)

var (
	DontReportExceptions []contracts.Exception
)

func init() {
	DontReportExceptions = []contracts.Exception{}
}

type ExceptionHandler struct {
	exceptions.DefaultExceptionHandler
}

func (handler ExceptionHandler) Handle(exception contracts.Exception) {
	if httpException, isHttpException := exception.(http.HttpException); isHttpException {
		fmt.Println(
			http.JsonResponse(contracts.Fields{
				"error":  exception.Error(),
				"fields": exception.GetFields(),
			}).Response(httpException.Context),
		)
		return
	}

	handler.DefaultExceptionHandler.Handle(exception)
}
