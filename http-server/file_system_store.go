package httpserver

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database io.Writer
	league   league
}

type league []Player

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, io.SeekStart)
	league, _ := NewLeague(file)

	return &FileSystemPlayerStore{
		database: &tape{file},
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() league {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	player := f.league.Find(playerName)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(playerName string) {
	player := f.league.Find(playerName)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{playerName, 1})
	}

	json.NewEncoder(f.database).Encode(f.league)
}
