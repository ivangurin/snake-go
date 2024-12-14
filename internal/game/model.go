package game

import (
	"image/color"
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

	font12 *text.GoTextFace
	font24 *text.GoTextFace

	colorSnakeHead = color.RGBA{0, 160, 100, 0}
	colorSnakeBody = color.RGBA{0, 220, 100, 0}
	colorFood      = color.RGBA{217, 30, 24, 0}
	colorText      = color.White
)
