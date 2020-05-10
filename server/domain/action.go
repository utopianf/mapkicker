package domain

type Action struct {
	seq           int
	participantID int
	kind          string
	mapIDs        []int
	msg           string
}

func NewChat(seq, participantID int, msg string) Action {
	return Action{
		seq:           seq,
		participantID: participantID,
		msg:           msg,
	}
}
