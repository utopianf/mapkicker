package usecase

import (
	"fmt"
	"mapkicker/domain"

	"github.com/gorilla/websocket"
)

type Socket struct {
	s               *websocket.Conn
	action          chan<- domain.Action
	broadcast       <-chan domain.Broadcast
	validationError <-chan domain.ValidationError
	close           chan string
}

func NewSocket(conn *websocket.Conn) domain.MessageChannel {
	socket := Socket{
		s:               conn,
		action:          make(chan domain.Action),
		broadcast:       make(chan domain.Broadcast),
		validationError: make(chan domain.ValidationError),
	}
	go socket.listen()
	return &socket
}

func (s *Socket) listen() {
	for {
		select {
		case b := <-s.broadcast:
			fmt.Println(b)
		case e := <-s.validationError:
			fmt.Println(e)
		case <-s.close:
			break
		}
	}
}

func (s *Socket) Action() chan<- domain.Action {
	return s.action
}

func (s *Socket) Broadcast() <-chan domain.Broadcast {
	return s.broadcast
}

func (s *Socket) ValidationError() <-chan domain.ValidationError {
	return s.validationError
}

const CLOSE = "close"

func (s *Socket) Close() error {
	s.close <- CLOSE
	return nil
}
