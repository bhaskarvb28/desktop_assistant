package state

import (
	"sync"
	"log"
)

type Manager struct {
	mu sync.RWMutex

	current AssistantState
}

func NewManager() *Manager {

	return &Manager{
		current: Idle,
	}
}

func (m *Manager) Get() AssistantState {

	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.current
}

func (m *Manager) Set(
	state AssistantState,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	log.Printf(
		"[STATE] %s → %s",
		m.current,
		state,
	)

	m.current = state
}