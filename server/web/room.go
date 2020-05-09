package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan []byte  // 転送すべきmessageを受け付けるchannel
	join    chan *client // これから参加するclientのためのchannel
	leave   chan *client // これから退室するclientのためのchannel
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join: // r.join <- (data) <sender>
			fmt.Println("Joinning")
			r.clients[client] = true
		case client := <-r.leave:
			fmt.Println("Leaving")
			r.clients[client] = false
			close(client.send)
		case msg := <-r.forward:
			fmt.Println("Message", msg)
			for client := range r.clients {
				select {
				case client.send <- msg:
					// submit message
				default: // 送信失敗した場合のエラー処理でclientを消す
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("ACCESS to /room from", req.Body)
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	fmt.Println("upgrade successful")
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() {
		r.leave <- client
		fmt.Println(fmt.Sprintf("client %v left", client))
	}()
	fmt.Println("HIHI")
	go client.write()
	client.read()
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}
