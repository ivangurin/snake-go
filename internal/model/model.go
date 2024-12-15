package model

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Point struct {
	X, Y int
}

var (
	DirectionUp    = Point{X: 0, Y: -1}
	DirectionDown  = Point{X: 0, Y: 1}
	DirectionLeft  = Point{X: -1, Y: 0}
	DirectionRight = Point{X: 1, Y: 0}

	Font12 *text.GoTextFace
	Font24 *text.GoTextFace

	ColorSnakeHead1 = color.RGBA{0, 128, 0, 0}
	ColorSnakeBody1 = color.RGBA{0, 225, 0, 0}
	ColorSnakeHead2 = color.RGBA{128, 0, 128, 0}
	ColorSnakeBody2 = color.RGBA{255, 0, 255, 0}
	ColorFood       = color.RGBA{255, 0, 0, 0}
	ColorText       = color.White
)
