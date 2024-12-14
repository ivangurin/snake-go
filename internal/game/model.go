package game

import (
	"snake-go/internal/food"
	"snake-go/internal/snake"
	"time"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	width      int
	height     int
	gridSize   int
	gameSpeed  time.Duration
	snake      *snake.Snake
	food       *food.Food
	lastUpdate time.Time
	gameOver   bool
}

var (
	textSource *text.GoTextFaceSource
	font12     *text.GoTextFace
	font24     *text.GoTextFace
)
