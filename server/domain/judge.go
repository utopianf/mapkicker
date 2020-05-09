package domain

// Judge は、1つのマップキックセッションの進行役を表す。
type Judge struct {
	id           int
	participants []Participant
}

// NewJudge は、指定したidを持つJudgeインスタンスを生成する。
// このidはchallongeトーナメントの試合IDを用いるので、外部から指定されるべきである。
func NewJudge(id int) *Judge {
	return &Judge{
		id: id,
	}
}

// Participants は、Judgeに属する参加者のスライスを返す。
func (j Judge) Participants() []Participant {
	return []Participant{}
}

// ID は、Judgeのidを返す。
func (j Judge) ID() int {
	return j.id
}
