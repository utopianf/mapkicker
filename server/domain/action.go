package domain

type Action struct {
	seq           int
	participantID int
	kind          string
	mapIDs        []int
	msg           string
}

type Broadcast struct {
	seq     int
	actions []Action
}

type ValidationError struct {
}

func NewChat(seq, participantID int, msg string) Action {
	return Action{
		seq:           seq,
		participantID: participantID,
		msg:           msg,
	}
}
