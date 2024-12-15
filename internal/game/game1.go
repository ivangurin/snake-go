package game

import (
	"fmt"
	"image/color"
	"snake-go/internal/food"
	"snake-go/internal/model"
	"snake-go/internal/snake"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type game1 struct {
	width    int
	height   int
	gridSize int
	snake    *snake.Snake
	food     *food.Food
	over     bool
}

func NewGame1(width, height, gridSize int) IGame {
	return &game1{
		width:    width,
		height:   height,
		gridSize: gridSize,
		snake:    snake.NewSnake(width, height, width/2, height/2, &model.DirectionRight),
		food:     food.NewFood(rand.Intn(width), rand.Intn(height)),
	}
}

func (g *game1) Update(move bool) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.snake.SetDirection(&model.DirectionUp)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.snake.SetDirection(&model.DirectionDown)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.snake.SetDirection(&model.DirectionLeft)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.snake.SetDirection(&model.DirectionRight)
	}

	if !move {
		return nil
	}

	g.snake.Move()

	if g.isCollision() {
		g.over = true
	}

	if *g.snake.GetHeadPoint() == *g.food.GetPoint() {
		g.snake.GrowUp()
		g.food = food.NewFood(rand.Intn(g.width), rand.Intn(g.height))
	}

	return nil
}

func (g *game1) Draw(screen *ebiten.Image) {
	if g.over {
		g.drawGameOver(screen)
		return
	}

	g.drawScore(screen)
	g.drawSnake(screen)
	g.drawFood(screen)
}

func (g *game1) IsOver() bool {
	return g.over
}

func (g *game1) isCollision() bool {
	headPoint := g.snake.GetHeadPoint()
	for _, point := range g.snake.GetPoints()[1:] {
		if *headPoint == point {
			return true
		}
	}
	return false
}

func (g *game1) drawFood(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(g.food.GetPoint().X*g.gridSize), float32(g.food.GetPoint().Y*g.gridSize), float32(g.gridSize), float32(g.gridSize), model.ColorFood, true)
}

func (g *game1) drawSnake(screen *ebiten.Image) {
	var clr color.RGBA
	for i, point := range g.snake.GetPoints() {
		clr = model.ColorSnakeBody1
		if i == 0 {
			clr = model.ColorSnakeHead1
		}
		vector.DrawFilledRect(screen, float32(point.X*g.gridSize), float32(point.Y*g.gridSize), float32(g.gridSize), float32(g.gridSize), clr, true)
	}
}

func (g *game1) drawScore(screen *ebiten.Image) {
	score := fmt.Sprintf("%d", g.snake.GetScore())
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(g.gridSize), float64(g.gridSize))
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, score, model.Font12, options)
}

func (g *game1) drawGameOver(screen *ebiten.Image) {
	gameOverText := "Game Over"
	textWidth, textHeight := text.Measure(gameOverText, model.Font24, model.Font24.Size)
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(g.width*g.gridSize/2)-float64(textWidth/2), float64(g.height*g.gridSize/2)-float64(textHeight/2)-50)
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, gameOverText, model.Font24, options)

	scoreText := fmt.Sprintf("Your score: %d", g.snake.GetScore())
	textWidth, textHeight = text.Measure(scoreText, model.Font12, model.Font12.Size)
	options = &text.DrawOptions{}
	options.GeoM.Translate(float64(g.width*g.gridSize/2)-float64(textWidth/2), float64(g.height*g.gridSize/2)-float64(textHeight/2)-5)
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, scoreText, model.Font12, options)

	pressEnterText := "Press Enter to continue"
	textWidth, textHeight = text.Measure(pressEnterText, model.Font12, model.Font12.Size)
	options = &text.DrawOptions{}
	options.GeoM.Translate(float64(g.width*g.gridSize/2)-float64(textWidth/2), float64(g.height*g.gridSize/2)-float64(textHeight/2)+30)
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, pressEnterText, model.Font12, options)
}
