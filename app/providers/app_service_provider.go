package providers

import (
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/jobs"
	"github.com/goal-web/example/app/listeners"
	"github.com/golang-module/carbon/v2"
)

type AppServiceProvider struct {
}

func (this AppServiceProvider) Register(app contracts.Application) {
	app.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)

		dispatcher.Register("QUERY_EXECUTED", listeners.DebugQuery{})
	})

	app.Call(func(serializer contracts.ClassSerializer) {
		serializer.Register(jobs.DemoClass)
	})
}

func (this AppServiceProvider) Start() error {

	return nil
}

func (this AppServiceProvider) Stop() {
}
