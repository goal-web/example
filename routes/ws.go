package routes

import (
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/websocket"
)

type DemoWSController struct {
}

func (d DemoWSController) OnConnect(request contracts.HttpRequest) error {
	if request.GetString("username") == "goal" {
		return nil
	}
	return errors.New("不允许连接")
}

func (d DemoWSController) OnMessage(frame contracts.WebSocketFrame) {
	fmt.Println("收到消息", frame.RawString(), frame.Connection().Fd())
	_ = frame.Send("来自服务器的回复1")
	_ = frame.SendBytes([]byte("来自服务器的回复2"))
}

func WebSocketRoutes(router contracts.Router) {
	router.Static("/", "/")

	router.Get("/ws", websocket.New(DemoWSController{}))

	router.Get("/ws", websocket.Default(func(frame contracts.WebSocketFrame) {

		fmt.Println("收到消息", frame.RawString(), frame.Connection().Fd())
		_ = frame.Send("来自服务器的回复1")
		_ = frame.SendBytes([]byte("来自服务器的回复2"))

	}))

}
