package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (fs *FileSystemPlayerStore) GetLeague() []Player {
	fs.database.Seek(0, io.SeekStart)

	league, _ := NewLeague(fs.database)
	return league
}

func (fs *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range fs.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}

func (fs *FileSystemPlayerStore) RecordWin(name string) {
	league := fs.GetLeague()

	for i, player := range league {
		if player.Name == name {
			player.Wins++
			league[i].Wins++
		}
	}

	fs.database.Seek(0, io.SeekStart)
	json.NewEncoder(fs.database).Encode(league)
}
