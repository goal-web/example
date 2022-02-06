package jobs

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/queue"
	"github.com/goal-web/supports/class"
	"github.com/goal-web/supports/logs"
)

var DemoClass = class.Make(new(Demo))

type Demo struct {
	*queue.Job
	Info string `json:"info"`
}

func NewDemo(info string) contracts.Job {
	return &Demo{
		Job:  queue.Base("default"),
		Info: info,
	}
}

func (demo *Demo) Handle() {
	logs.Default().WithField("info", demo.Info).Info("demo job")
}
