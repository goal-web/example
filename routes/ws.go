package routes

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/http"
	"github.com/gorilla/websocket"
)

var (
	upgrade = websocket.Upgrader{}
)

func WebSocketRoutes(router contracts.Router) {

	router.Static("/", "/")

	router.Get("/ws", WebSocket)
}

func WebSocket(request *http.Request) error {
	ws, err := upgrade.Upgrade(request.Context.Response(), request.Request(), nil)
	if err != nil {
		return err
	}

	defer ws.Close()

	for {
		// Write
		err = ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			request.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			request.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
