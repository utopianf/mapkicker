package domain_test

import (
	"mapkicker/domain"
	"testing"
)

func TestNewJudge(t *testing.T) {
	judgeID := 1
	j := domain.NewJudge(judgeID)
	if j.ID() != judgeID {
		t.Errorf("judge is not properly initialized. expect %v, actual %v", judgeID, j.ID())
	}
}
