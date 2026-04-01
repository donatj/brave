package main

import (
	"os"

	"github.com/donatj/brave/game"
)

func main() {
	// Create a new game instance with default settings
	g := game.NewGame(os.Stdin, os.Stdout)

	// Run the game
	g.Run()
}
