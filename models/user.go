package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var (
	UserModel = table.NewModel(class.Make(new(User)), "users")
)

func UserQuery() *table.Table {
	return table.FromModel(UserModel)
}

type User struct {
	Id       string `json:"id"`
	NickName string `json:"name"`
}

func (u User) GetId() string {
	return u.Id
}
