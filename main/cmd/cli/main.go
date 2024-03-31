package main

import (
	"fmt"
	poker "hello"
	"os"
)

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	store := poker.NewInMemoryPlayerStore()
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
