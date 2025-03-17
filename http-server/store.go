package main

import "sync"

// PlayerStore defines methods to get a player's score and record wins.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	mu    sync.RWMutex
	store map[string]int
}

// NewInMemoryPlayerStore initialises an empty player store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
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
