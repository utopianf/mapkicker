package repository

import (
	"mapkicker/domain"
)

// OnMemory は、dbに依存しない揮発性のrepository.
type OnMemory struct {
	manager *domain.Manager
}

// Manager は、Managerインスタンスを返す。
func (r *OnMemory) Manager() *domain.Manager {
	return r.manager
}

// AllMaps は、現在のマッププールに属するすべてのマップを返す
func (r *OnMemory) AllMaps() []domain.Map {
	return nil
}

// NewRepository は、リポジトリを得る。
func NewRepository() domain.Repository {
	return &OnMemory{
		manager: domain.GetManager(),
	}
}
