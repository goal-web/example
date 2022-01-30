package controllers

import (
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/table"
	"github.com/goal-web/example/models"
	"github.com/goal-web/supports/utils"
	"strconv"
)

func HelloWorld() string {
	return "hello, goal."
}

func Counter(session contracts.Session) string {
	count := utils.ConvertToInt(session.Get("count", "0"), 0)
	count++
	session.Put("count", strconv.Itoa(count))
	panic(errors.New("报错"))
	return "hello, goal." + strconv.Itoa(count)
}

func DatabaseQuery(db contracts.DBFactory, request contracts.HttpRequest) contracts.Fields {
	connection := db.Connection(request.GetString("connection"))
	var user struct {
		Id   int    `db:"id"`
		Name string `db:"name"`
	}
	err := connection.Get(&user, fmt.Sprintf("select * from users where name='%s'", "qbhy"))
	return contracts.Fields{
		"user": user,
		"err":  err,
	}
}

func RedisExample(redis contracts.RedisConnection) contracts.Fields {
	str, err := redis.Get("incr")
	return contracts.Fields{
		"value": str,
		"err":   err,
	}
}

func GetUsers() interface{} {
	models.UserModel().Create(map[string]interface{}{
		"name": utils.RandStr(10),
	})
	return contracts.Fields{
		"models": models.UserModel().Get().ToArrayFields(),
		"tables": table.Query("users").Get().ToArrayFields(),
	}
}
