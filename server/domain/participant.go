package domain

import "github.com/gorilla/websocket"

// Participant は、Mapkickへの参加者を表す。
type Participant struct {
	socket  *websocket.Conn
	msgChan MessageChannel
	j       *Judge
}

func NewParticipant(socket *websocket.Conn) *Participant {
	return &Participant{
		socket: socket,
	}
}
