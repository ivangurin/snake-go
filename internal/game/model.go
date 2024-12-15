package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type IGame interface {
	Draw(screen *ebiten.Image)
	Update(move bool) error
	IsOver() bool
}
