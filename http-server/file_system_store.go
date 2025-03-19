package main

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore stores players in the filesystem.
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)

	return &FileSystemPlayerStore{
		database: database,
		league:   league,
	}
}

// GetLeague returns the scores of all the players.
func (fs *FileSystemPlayerStore) GetLeague() League {
	return fs.league
}

// GetPlayerScore retrieves a player's score.
func (fs *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := fs.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin will store a win for a player, incrementing wins if already known.
func (fs *FileSystemPlayerStore) RecordWin(name string) {
	player := fs.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		fs.league = append(fs.league, Player{name, 1})
	}

	fs.database.Seek(0, io.SeekStart)
	json.NewEncoder(fs.database).Encode(fs.league)
}
