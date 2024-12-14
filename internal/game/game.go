package game

import (
	"bytes"
	"fmt"
	"image/color"
	"math/rand"
	"snake-go/internal/food"
	"snake-go/internal/model"
	"snake-go/internal/snake"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func NewGame(width, height int) (*Game, error) {
	var err error
	textSource, err = text.NewGoTextFaceSource(
		bytes.NewReader(
			fonts.PressStart2P_ttf,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load font: %w", err)
	}

	font12 = &text.GoTextFace{
		Source: textSource,
		Size:   12,
	}

	font24 = &text.GoTextFace{
		Source: textSource,
		Size:   24,
	}

	return &Game{
		width:     width,
		height:    height,
		gridSize:  20,
		gameSpeed: time.Second / 6,
	}, nil
}

func (g *Game) Run() error {
	ebiten.SetWindowSize(g.width*g.gridSize, g.height*g.gridSize)
	ebiten.SetWindowTitle("Snake Game")
	g.startNewGame()
	err := ebiten.RunGame(g)
	if err != nil {
		return fmt.Errorf("failed to run game: %w", err)
	}
	return nil
}

func (g *Game) startNewGame() {
	g.snake = snake.NewSnake(g.width, g.height, &model.DirectionRight)
	g.food = food.NewFood(rand.Intn(g.width), rand.Intn(g.height))
	g.gameOver = false
}

func (g *Game) stopGame() {
	g.gameOver = true
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.startNewGame()
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.snake.SetDirection(&model.DirectionUp)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.snake.SetDirection(&model.DirectionDown)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.snake.SetDirection(&model.DirectionLeft)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.snake.SetDirection(&model.DirectionRight)
	}

	if time.Since(g.lastUpdate) < g.gameSpeed {
		return nil
	}

	g.snake.Move()

	if g.isCollision() {
		g.stopGame()
		return nil
	}

	if *g.snake.GetHeadPoint() == *g.food.GetPoint() {
		g.snake.GrowUp()
		g.food = food.NewFood(rand.Intn(g.width), rand.Intn(g.height))
	}

	g.lastUpdate = time.Now()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameOver {
		g.drawGameOver(screen)
		return
	}

	g.drawScore(screen)
	g.drawSnake(screen, g.snake)
	g.drawFood(screen, g.food)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width * g.gridSize, g.height * g.gridSize
}

func (g *Game) isCollision() bool {
	headPoint := g.snake.GetHeadPoint()
	for _, point := range g.snake.GetPoints()[1:] {
		if *headPoint == point {
			return true
		}
	}
	return false
}

func (g *Game) drawSnake(screen *ebiten.Image, snake *snake.Snake) {
	var clr color.RGBA
	for i, point := range snake.GetPoints() {
		clr = colorSnakeBody
		if i == 0 {
			clr = colorSnakeHead
		}
		vector.DrawFilledRect(screen, float32(point.X*g.gridSize), float32(point.Y*g.gridSize), float32(g.gridSize), float32(g.gridSize), clr, true)
	}
}

func (g *Game) drawFood(screen *ebiten.Image, food *food.Food) {
	vector.DrawFilledRect(screen, float32(food.GetPoint().X*g.gridSize), float32(food.GetPoint().Y*g.gridSize), float32(g.gridSize), float32(g.gridSize), colorFood, true)
}

func (g *Game) drawScore(screen *ebiten.Image) {
	score := fmt.Sprintf("%d", g.snake.GetScore())
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(g.gridSize), float64(g.gridSize))
	options.ColorScale.ScaleWithColor(colorText)
	text.Draw(screen, score, font12, options)
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	gameOverText := "Game Over"
	textWidth, textHeight := text.Measure(gameOverText, font24, font24.Size)
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(g.width*g.gridSize/2)-float64(textWidth/2), float64(g.height*g.gridSize/2)-float64(textHeight/2)-30)
	options.ColorScale.ScaleWithColor(colorText)
	text.Draw(screen, gameOverText, font24, options)

	pressEnterText := "Press Enter to start again"
	textWidth, textHeight = text.Measure(pressEnterText, font12, font12.Size)
	options = &text.DrawOptions{}
	options.GeoM.Translate(float64(g.width*g.gridSize/2)-float64(textWidth/2), float64(g.height*g.gridSize/2)-float64(textHeight/2)+10)
	options.ColorScale.ScaleWithColor(colorText)
	text.Draw(screen, pressEnterText, font12, options)
}
