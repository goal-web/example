package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/jobs"
	"github.com/golang-module/carbon/v2"
	"time"
)

func DemoJob(queue contracts.Queue, request contracts.HttpRequest) string {
	queue.Push(jobs.NewDemo(request.GetString("info")))
	queue.Later(
		time.Now().Add(time.Second*time.Duration(request.GetInt("delay"))),
		jobs.NewDemo("delay+"+request.GetString("info")),
	)

	return carbon.Now().String()
}
