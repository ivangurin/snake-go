package main

import (
	game "snake-go/internal/app"
)

func main() {
	app, err := game.NewApp(30, 20)
	if err != nil {
		panic(err)
	}

	app.Run()
}
