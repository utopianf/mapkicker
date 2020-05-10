package domain

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Judge は、1つのマップキックセッションの進行役を表す。
type Judge struct {
	id           int
	join         chan *Participant
	participants []*Participant
	history      []string
}

// NewJudge は、指定したidを持つJudgeインスタンスを生成する。
// このidはchallongeトーナメントの試合IDを用いるので、外部から指定されるべきである。
func NewJudge(id int) *Judge {
	return &Judge{
		id:   id,
		join: make(chan *Participant),
	}
}

func (j *Judge) AddNewParticipant(socket *websocket.Conn) *Participant {
	p := NewParticipant(socket)
	fmt.Println("trying to join")
	j.join <- p
	fmt.Printf("Judge has %v participants now!\n", len(j.participants))
	for _, msg := range j.history {
		p.socket.WriteMessage(websocket.TextMessage, []byte(msg))
	}
	return p
}

func (j *Judge) Run() {
	go func() {
		for p := range j.join {
			j.participants = append(j.participants, p)
		}
	}()
	fmt.Println("Judge is running")
}

func (j *Judge) Share(msg string) {
	j.history = append(j.history, msg)
	for _, p := range j.participants {
		p.socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%v", msg)))
	}
}

// Participants は、Judgeに属する参加者のスライスを返す。
func (j *Judge) Participants() []Participant {
	return []Participant{}
}

// ID は、Judgeのidを返す。
func (j *Judge) ID() int {
	return j.id
}
