package domain

// Manager は、複数のJudgeを取りまとめる。Application全体で1つしか生成されない。
type Manager struct {
	judges map[int]*Judge
}

// GetManager は、Managerのインスタンスを得る。
func GetManager() *Manager {
	return &Manager{
		judges: make(map[int]*Judge),
	}
}
