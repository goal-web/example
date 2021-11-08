package controllers

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/validation"
	"goal-example/app/requests"
	"time"
)

// Hashing 测试哈希操作
func Hashing(hash contracts.HasherFactory) string {
	res := fmt.Sprintf("bcrypt 哈希值：%s\n", hash.Make("哈希", nil))
	res += fmt.Sprintf("md5 哈希值：%s", hash.Driver("md5").Make("哈希", nil))

	return res
}

// Redis 测试 redis 操作
func Redis(redis contracts.RedisFactory) {
	fmt.Println(redis.Connection("cache").Get("a"))
}

// Filesystem 测试文件系统操作
func Filesystem(filesystem contracts.FileSystemFactory) string {
	contents, err := filesystem.Get(".env")
	if err != nil {
		panic(exceptions.ResolveException(err))
	}

	return contents
}

// Cache 测试缓存操作
func Cache(cache contracts.CacheStore) string {
	err := cache.Put("goal", "goal", time.Hour)
	if err != nil {
		panic(exceptions.ResolveException(err))
	}

	return cache.Get("goal").(string)
}

// DiRequest 测试注入自定义请求实例
func DiRequest(request requests.LoginRequest) interface{} {
	result := validation.Valid(&request)
	//result.Validate() // 这个方法如果数据校验不通过，则直接抛异常
	if result.IsFail() {
		return result.Errors()
	}
	return fmt.Sprintf("账号是：%s，密码是：%s", request.Get("name"), request.Get("password"))
}
