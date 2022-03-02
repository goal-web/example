package routes

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/websocket"
)

func WebSocket(router contracts.Router) {
	router.Static("/", "/")

	router.Get("/ws", websocket.Default(func(frame contracts.WebSocketFrame) {

		fmt.Println("收到消息", frame.RawString(), frame.Connection().Fd())
		_ = frame.Send("来自服务器的回复1")
		_ = frame.SendBytes([]byte("来自服务器的回复2"))

	}))

}
