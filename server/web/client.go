package web

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}

func (c *client) read() { // clientがメッセージを送信したときに呼ばれる 主語がserver
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			fmt.Println(fmt.Sprintf("client.read %v", msg))
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}
func (c *client) write() { // clientにサーバーからメッセージを配信するときに使う 主語がserver
	for msg := range c.send {
		fmt.Println(fmt.Sprintf("client.write %v", msg))
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
