package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/email"
	"testing"
)

func TestSendEmail(t *testing.T) {
	app := initApp()

	app.Call(func(mailer contracts.Mailer, config contracts.Config) {
		fmt.Println(config.Get("mail"))
		fmt.Println(mailer.Send(email.New("测试邮件", email.Text("")).SetTo("572490755@qq.com")))
	})
}
