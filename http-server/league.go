package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

func (l league) Find(playerName string) *Player {
	for i, p := range l {
		if p.Name == playerName {
			return &l[i]
		}
	}
	return nil
}
