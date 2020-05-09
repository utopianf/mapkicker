package domain

// Repository は、リポジトリのインタフェース
type Repository interface {
	Manager() *Manager
	AllMaps() []Map
}
