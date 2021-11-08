package requests

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/http"
	"github.com/qbhy/goal/validation/checkers"
)

type LoginRequest struct {
	http.Request `di:""` // di 注解表示需要注入
}

func (this *LoginRequest) GetFields() contracts.Fields {
	return this.All()
}

func (this *LoginRequest) Rules() contracts.Rules {
	return contracts.Rules{
		"name": []contracts.Checker{
			checkers.StrLen(2, 6), // 字符串长度在 2-6 之间
		},
		"password": []contracts.Checker{
			checkers.StrLen(6, 16), // 字符串长度在 6-16 之间
		},
	}
}
