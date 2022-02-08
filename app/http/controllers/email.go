package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/email"
)

func SendEmail(request contracts.HttpRequest, mailer contracts.Mailer) string {
	err := mailer.Send(
		email.New("测试邮件", email.Text(request.GetString("content"))).
			SetTo(request.GetString("to")),
	)

	if err != nil {
		return err.Error()
	}

	return "ok"
}
