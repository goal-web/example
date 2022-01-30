package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var UserClass = class.Make(new(User))

func UserModel() *table.Table {
	return table.Model(UserClass, "users")
}

type User struct {
	Id       int64  `json:"id"`
	NickName string `json:"name"`
}
