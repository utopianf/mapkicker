package domain

import "io"

// MessageChannel は、Participantが依存するWebSocketなどの通信インフラを抽象化したもの
// 実装はusecaseかwebにおかれる(多分usecase)
type MessageChannel interface {
	Action() chan<- Action
	Broadcast() <-chan Broadcast
	ValidationError() <-chan ValidationError
	io.Closer
}
