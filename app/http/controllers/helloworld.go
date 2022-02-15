package controllers

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/table"
	"github.com/goal-web/example/app/models"
	"github.com/goal-web/supports/utils"
	"strconv"
)

/**
Running 30s test @ http://localhost:8008
  10 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    21.65ms   58.26ms   1.36s    91.14%
    Req/Sec     2.39k     2.16k   11.90k    68.85%
  615610 requests in 30.09s, 75.73MB read
  Socket errors: connect 759, read 34, write 0, timeout 0
Requests/sec:  20456.33
Transfer/sec:      2.52MB
*/
func HelloWorld() string {
	return "hello, goal."
}

func Counter(session contracts.Session) string {
	count := utils.ConvertToInt(session.Get("count", "0"), 0)
	count++
	session.Put("count", strconv.Itoa(count))
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
	models.UserQuery().Create(map[string]interface{}{
		"name": utils.RandStr(10),
	})
	return contracts.Fields{
		"models": models.UserQuery().Get().ToArrayFields(),
		"tables": table.Query("users").Get().ToArrayFields(),
	}
}
