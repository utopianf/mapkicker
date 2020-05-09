package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	go func() {
		for {
			if _, msg, err := socket.ReadMessage(); err == nil {
				fmt.Println(fmt.Sprintf("Read msg %v", msg))
				socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ECHO: %v", msg)))
			} else {
				break
			}
		}
	}()
}
