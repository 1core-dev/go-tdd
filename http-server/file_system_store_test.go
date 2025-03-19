package main

import (
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[
		{"Name": "John", "Wins": 42},
		{"Name": "Sean", "Wins": 24}]`)
	defer cleanDatabase()
	store := NewFileSystemPlayerStore(database)

	t.Run("league from reader", func(t *testing.T) {

		got := store.GetLeague()

		want := []Player{
			{"John", 42},
			{"Sean", 24},
		}
		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
	t.Run("get player score", func(t *testing.T) {
		got := store.GetPlayerScore("Sean")
		want := 24
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("John")

		got := store.GetPlayerScore("John")
		want := 43
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for new player", func(t *testing.T) {
		playerName := "Peter"

		store.RecordWin(playerName)

		got := store.GetPlayerScore(playerName)
		want := 1
		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t *testing.T, InitData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Errorf("could not create temp file %v", err)
	}

	tmpFile.Write([]byte(InitData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
