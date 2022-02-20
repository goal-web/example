package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/table"
	"github.com/goal-web/example/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getQuery(name string) contracts.QueryBuilder {
	// 测试用例环境下的简易 goal 应用启动
	app := initApp()

	//return  table.Query("users") 返回 table 实例，使用默认连接
	//tx, _ := app.Get("db").(contracts.DBConnection).Begin()
	//return table.WithTX("users", tx) // 事物环境下执行

	//return table.WithConnection(name, "sqlite") // 返回指定连接的 table 实例，使用连接名
	return table.WithConnection(name, app.Get("db").(contracts.DBConnection)) // 也可以指定连接实例
}

// TestTableQuery 测试不带模型的 table 查询，类似 laravel 的 DB::table()
func TestTableQuery(t *testing.T) {

	getQuery("users").Delete()

	// 不设置模型的情况下，返回 contracts.Fields
	user := getQuery("users").Create(contracts.Fields{
		"name": "qbhy",
	}).(contracts.Fields)

	fmt.Println(user)
	userId := user["id"].(int64)
	// 判断插入是否成功
	assert.True(t, userId > 0)

	// 获取数据总量
	assert.True(t, getQuery("users").Count() == 1)

	// 修改数据
	num := getQuery("users").Where("name", "qbhy").Update(contracts.Fields{
		"name": "goal",
	})
	assert.True(t, num == 1)
	// 判断修改后的数据
	user = getQuery("users").Where("name", "goal").First().(contracts.Fields)
	fmt.Println("user", user)

	err := getQuery("users").Chunk(10, func(collection contracts.Collection, page int) error {
		assert.True(t, collection.Len() == 1)
		fmt.Println(collection.ToJson())
		return nil
	})

	assert.Nil(t, err)

	assert.True(t, user["id"] == userId)
	assert.True(t, user["name"] == "goal")
	assert.True(t, getQuery("users").Find(userId).(contracts.Fields)["id"] == userId)
	fmt.Println("find:", getQuery("users").Find(userId))
	assert.True(t, getQuery("users").Where("id", userId).Delete() == 1)
	assert.Nil(t, getQuery("users").Find(userId))
}

func TestModel(t *testing.T) {
	initApp()

	user := models.UserQuery().Create(contracts.Fields{
		"name": "qbhy",
	}).(models.User)

	fmt.Println("创建后返回模型", user)

	fmt.Println("用table查询：",
		getQuery("users").Get().Map(func(user contracts.Fields) {
			fmt.Println("用table查询", user)
		}).ToJson()) // query 返回 Collection<contracts.Fields>

	fmt.Println(models.UserQuery(). // model 返回 Collection<models.User>
					Get().
					Map(func(user models.User) {
			fmt.Println("id:", user.GetId())
		}).ToJson())

	fmt.Println(models.UserQuery().Where("id", ">", 0).Delete())
}

//func TestClickhouse(t *testing.T) {
//	initApp()
//
//	var UserQuery = func() *table.Table {
//		return table.WithConnection("users", "clickhouse")
//	}
//
//	user := UserQuery().Create(contracts.Fields{
//		"name": "qbhy",
//	})
//
//	fmt.Println("创建后返回", user)
//
//	fmt.Println("用table查询：",
//		UserQuery().Get().Map(func(user contracts.Fields) {
//			fmt.Println("用table查询", user)
//		}).ToJson()) // query 返回 Collection<contracts.Fields>
//
//	fmt.Println(UserQuery().Where("id", ">", 0).Delete())
//}
