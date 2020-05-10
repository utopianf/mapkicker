package domain_test

import (
	"mapkicker/domain"
	"testing"
)

func TestNewChat(t *testing.T) {
	j := domain.NewJudge(0)
	action := domain.NewChat(0, 1, "hello")
	ok := j.Send(action)
	if !ok {
		t.Error("New judge seq should began with 0")
	}
	ok = j.Send(action)
	if ok {
		t.Error("judge seq is not correctly incremented")
	}
}
