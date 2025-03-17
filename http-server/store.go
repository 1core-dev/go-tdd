package main

import "sync"

// NewInMemoryPlayerStore initialises an empty player store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
}

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	mu    sync.RWMutex
	store map[string]int
}

// RecordWin will record a player's win.
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.store[name]++
}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.RLock()
	defer i.mu.RUnlock()

	return i.store[name]
}
