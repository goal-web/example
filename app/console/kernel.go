package console

import (
	"github.com/goal-web/application/commands"
	"github.com/goal-web/console"
	"github.com/goal-web/contracts"
	commands2 "github.com/goal-web/example/app/console/commands"
)

func NewKernel(app contracts.Application) contracts.Console {
	return &Kernel{console.NewKernel(app, []contracts.CommandProvider{
		commands.Runner,
		commands2.NewHello,
	}), app}
}

type Kernel struct {
	*console.Kernel
	app contracts.Application
}

func (this *Kernel) Schedule(schedule contracts.Schedule) {
	//schedule.Call(func() {
	//	logs.Default().Info("每10秒钟打印 goal")
	//}).EveryTenSeconds()
	//
	//schedule.Call(func() {
	//	time.Sleep(time.Second * 3)
	//	logs.Default().Info("每隔五秒打印by WithoutOverlapping")
	//}).Description("每隔五秒打印by WithoutOverlapping").
	//	WithoutOverlapping(10).
	//	EverySecond()
	//
	//schedule.Call(func() {
	//	logs.Default().Info("八点到九点，每秒更新一次")
	//}).EverySecond().Between("20:00", "22:00")
	//
	//schedule.Command(commands2.NewHello(this.app), "每秒钟").EverySecond().Between("20:00", "23:59")
	//
	//schedule.Exec("hello", "隔五秒").EveryFiveSeconds().Between("20:00", "23:59")
	//
	//schedule.Call(func() {
	//	logs.Default().Info("周日每5秒钟打印 周日愉快")
	//}).EveryFiveSeconds().Sundays()
}
