package app

import (
	"snake-go/internal/game"
	"time"
)

const (
	modeMenu = iota
	modeGame
)

type app struct {
	width      int
	height     int
	gridSize   int
	gameSpeed  time.Duration
	lastUpdate time.Time
	game       game.IGame
	mode       int
}
