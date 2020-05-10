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
		// N回echoしたらsocketを閉じてみる
		N := 5
		for i := 0; i < N; i++ {
			if _, msg, err := socket.ReadMessage(); err == nil {
				s := string(msg)
				fmt.Println(fmt.Sprintf("Read msg %v", s))
				socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ECHO: %v", s)))
			} else {
				break
			}
		}
		if err := socket.Close(); err != nil {
			fmt.Println(err)
		}
	}()
}
