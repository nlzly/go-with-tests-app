package httpserver

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

type league []Player

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	store := new(FileSystemPlayerStore)
	store.database = database
	return store
}

func (f *FileSystemPlayerStore) GetLeague() league {
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	player := f.GetLeague().Find(playerName)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(playerName string) {
	league := f.GetLeague()
	player := league.Find(playerName)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{playerName, 1})
	}

	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}

func (l league) Find(playerName string) *Player {
	for i, p := range l {
		if p.Name == playerName {
			return &l[i]
		}
	}
	return nil
}
