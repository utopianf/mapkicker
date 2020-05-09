package repository

import (
	"mapkicker/domain"
)

type RepositoryOnMemory struct {
	manager *domain.Manager
}

func (r *RepositoryOnMemory) RegisterManager(m *domain.Manager) error {
	r.manager = m
	return nil
}

func (r *RepositoryOnMemory) Manager() *domain.Manager {
	return r.manager
}

// NewRepository は、リポジトリを得る。
func NewRepository() domain.Repository {
	return &RepositoryOnMemory{
		manager: domain.GetManager(),
	}
}
