package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/jobs"
)

func DemoJob(queue contracts.Queue, request contracts.HttpRequest) string {
	queue.Push(jobs.NewDemo(request.GetString("info")))

	return "ok"
}
