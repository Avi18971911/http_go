package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in: bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	for {
		cli.in.Scan()
		cli.playerStore.RecordWin(extractWinner(cli.in.Text()))
	}
}

func extractWinner(inputText string) string {
	return strings.Replace(inputText, " wins", "", 1)
}
