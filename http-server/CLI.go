package httpserver

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	reader      io.Reader
}

func NewCLI(store PlayerStore, reader io.Reader) *CLI {
	cli := new(CLI)
	cli.playerStore = store
	cli.reader = reader

	return cli
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.reader)
	reader.Scan()
	cli.playerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
