package domain

type Repository interface {
	RegisterManager(*Manager) error
	Manager() *Manager
}
