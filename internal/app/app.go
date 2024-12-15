package app

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"snake-go/internal/game"
	"snake-go/internal/model"
)

func NewApp(width, height int) (*app, error) {
	var err error
	textSource, err := text.NewGoTextFaceSource(
		bytes.NewReader(
			fonts.PressStart2P_ttf,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load font: %w", err)
	}

	model.Font12 = &text.GoTextFace{
		Source: textSource,
		Size:   12,
	}

	model.Font24 = &text.GoTextFace{
		Source: textSource,
		Size:   24,
	}

	return &app{
		width:     width,
		height:    height,
		gridSize:  20,
		gameSpeed: time.Second / 6,
		mode:      modeMenu,
	}, nil
}

func (a *app) Run() error {
	ebiten.SetWindowSize(a.width*a.gridSize, a.height*a.gridSize)
	ebiten.SetWindowTitle("Snake Game")
	err := ebiten.RunGame(a)
	if err != nil {
		return fmt.Errorf("failed to run game: %w", err)
	}
	return nil
}

func (a *app) startNewGame(players int) {
	if players == 1 {
		a.game = game.NewGame1(a.width, a.height, a.gridSize)
	} else {
		a.game = game.NewGame2(a.width, a.height, a.gridSize)
	}
	a.mode = modeGame
}

func (a *app) Update() error {
	if a.mode == modeMenu {
		if ebiten.IsKeyPressed(ebiten.Key1) {
			a.startNewGame(1)
		}
		if ebiten.IsKeyPressed(ebiten.Key2) {
			a.startNewGame(2)
		}
		return nil
	}

	if a.game.IsOver() {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			a.mode = modeMenu
		}
		return nil
	}

	var move bool
	if a.lastUpdate.IsZero() {
		a.lastUpdate = time.Now()
	}
	if time.Since(a.lastUpdate) > a.gameSpeed {
		move = true
		a.lastUpdate = time.Now()
	}

	err := a.game.Update(move)
	if err != nil {
		return fmt.Errorf("failed to update game: %w", err)
	}

	return nil
}

func (a *app) Draw(screen *ebiten.Image) {
	if a.mode == modeMenu {
		a.drawMenu(screen)
		return
	}

	a.game.Draw(screen)
}

func (a *app) Layout(outsideWidth, outsideHeight int) (int, int) {
	return a.width * a.gridSize, a.height * a.gridSize
}

func (a *app) drawMenu(screen *ebiten.Image) {
	gameOverText := "Snake Game"
	textWidth, textHeight := text.Measure(gameOverText, model.Font24, model.Font24.Size)
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(a.width*a.gridSize/2)-float64(textWidth/2), float64(a.height*a.gridSize/2)-float64(textHeight/2)-50)
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, gameOverText, model.Font24, options)

	textMenu1 := "Press 1 for 1 player"
	textWidth, textHeight = text.Measure(textMenu1, model.Font12, model.Font12.Size)
	options = &text.DrawOptions{}
	options.GeoM.Translate(float64(a.width*a.gridSize/2)-float64(textWidth/2), float64(a.height*a.gridSize/2)-float64(textHeight/2)-5)
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, textMenu1, model.Font12, options)

	pressEnterText := "Press 2 for 2 players"
	textWidth, textHeight = text.Measure(pressEnterText, model.Font12, model.Font12.Size)
	options = &text.DrawOptions{}
	options.GeoM.Translate(float64(a.width*a.gridSize/2)-float64(textWidth/2), float64(a.height*a.gridSize/2)-float64(textHeight/2)+30)
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, pressEnterText, model.Font12, options)
}
