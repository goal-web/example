package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
)

func LoginExample(guard contracts.Guard) contracts.Fields {
	//  这是伪代码
	user := models.UserQuery().First().(models.User)

	return contracts.Fields{
		"token": guard.Login(user), // jwt 返回 token，session 返回 true
	}
}

func GetCurrentUser(guard contracts.Guard) interface{} {
	return contracts.Fields{
		"user": guard.User(),
	}
}
