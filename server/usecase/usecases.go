package usecase

import (
	"fmt"
	"mapkicker/domain"
)

type CreateJudgeUseCase struct {
	r domain.Repository
}

func (c CreateJudgeUseCase) Create(id int) {
	j := domain.NewJudge(id)
	fmt.Println(j)
}
