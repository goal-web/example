package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/http/requests"
	"github.com/goal-web/example/app/models"
	"github.com/goal-web/validation"
)

func LoginExample(guard contracts.Guard, request requests.LoginRequest) contracts.Fields {

	validation.VerifyForm(request)

	//  这是伪代码
	var user, ok = models.UserQuery().Where("name", request.GetString("username")).First().(models.User)

	if !ok {
		return contracts.Fields{
			"error": "用户不存在",
		}
	}

	return contracts.Fields{
		"token": guard.Login(user), // jwt 返回 token，session 返回 true
	}
}

func GetCurrentUser(guard contracts.Guard) interface{} {
	return contracts.Fields{
		"user": guard.User(),
	}
}
