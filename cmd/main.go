package main

import (
	"snake-go/internal/game"
)

func main() {
	game, err := game.NewGame(
		30,
		20,
	)
	if err != nil {
		panic(err)
	}

	game.Run()
}
