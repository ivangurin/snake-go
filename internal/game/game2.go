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

type game2 struct {
	width    int
	height   int
	gridSize int
	snake1   *snake.Snake
	snake2   *snake.Snake
	food1    *food.Food
	food2    *food.Food
	over     bool
	winner   int
}

func NewGame2(width, height, gridSize int) IGame {
	return &game2{
		width:    width,
		height:   height,
		gridSize: gridSize,
		snake1:   snake.NewSnake(width, height, width/4, height/2, &model.DirectionRight),
		snake2:   snake.NewSnake(width, height, width-width/4, height/2, &model.DirectionLeft),
		food1:    food.NewFood(rand.Intn(width), rand.Intn(height)),
		food2:    food.NewFood(rand.Intn(width), rand.Intn(height)),
	}
}

func (g *game2) Update(move bool) error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.snake1.SetDirection(&model.DirectionUp)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.snake1.SetDirection(&model.DirectionDown)
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.snake1.SetDirection(&model.DirectionLeft)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.snake1.SetDirection(&model.DirectionRight)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.snake2.SetDirection(&model.DirectionUp)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.snake2.SetDirection(&model.DirectionDown)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.snake2.SetDirection(&model.DirectionLeft)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.snake2.SetDirection(&model.DirectionRight)
	}

	if !move {
		return nil
	}

	g.snake1.Move()
	g.snake2.Move()

	if g.isCollision() {
		g.over = true
		return nil
	}

	if *g.snake1.GetHeadPoint() == *g.food1.GetPoint() {
		g.snake1.GrowUp()
		g.food1 = food.NewFood(rand.Intn(g.width), rand.Intn(g.height))
	}

	if *g.snake1.GetHeadPoint() == *g.food2.GetPoint() {
		g.snake1.GrowUp()
		g.food2 = food.NewFood(rand.Intn(g.width), rand.Intn(g.height))
	}

	if *g.snake2.GetHeadPoint() == *g.food1.GetPoint() {
		g.snake2.GrowUp()
		g.food1 = food.NewFood(rand.Intn(g.width), rand.Intn(g.height))
	}

	if *g.snake2.GetHeadPoint() == *g.food2.GetPoint() {
		g.snake2.GrowUp()
		g.food2 = food.NewFood(rand.Intn(g.width), rand.Intn(g.height))
	}

	return nil
}

func (g *game2) Draw(screen *ebiten.Image) {
	if g.over {
		g.drawGameOver(screen)
		return
	}

	g.drawScore(screen)
	g.drawSnake(screen)
	g.drawFood(screen)
}

func (g *game2) isCollision() bool {
	headPoint1 := g.snake1.GetHeadPoint()
	headPoint2 := g.snake2.GetHeadPoint()

	if headPoint1 == headPoint2 {
		return true
	}

	for _, point := range g.snake1.GetPoints()[1:] {
		if *headPoint1 == point {
			g.winner = 2
			return true
		}
	}
	for _, point := range g.snake2.GetPoints() {
		if *headPoint1 == point {
			g.winner = 2
			return true
		}
	}

	for _, point := range g.snake2.GetPoints()[1:] {
		if *headPoint2 == point {
			g.winner = 1
			return true
		}
	}
	for _, point := range g.snake1.GetPoints() {
		if *headPoint2 == point {
			g.winner = 1
			return true
		}
	}

	return false
}

func (g *game2) IsOver() bool {
	return g.over
}

func (g *game2) drawFood(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(g.food1.GetPoint().X*g.gridSize), float32(g.food1.GetPoint().Y*g.gridSize), float32(g.gridSize), float32(g.gridSize), model.ColorFood, true)
	vector.DrawFilledRect(screen, float32(g.food2.GetPoint().X*g.gridSize), float32(g.food2.GetPoint().Y*g.gridSize), float32(g.gridSize), float32(g.gridSize), model.ColorFood, true)
}

func (g *game2) drawSnake(screen *ebiten.Image) {
	var clr color.RGBA
	for i, point := range g.snake1.GetPoints() {
		clr = model.ColorSnakeBody1
		if i == 0 {
			clr = model.ColorSnakeHead1
		}
		vector.DrawFilledRect(screen, float32(point.X*g.gridSize), float32(point.Y*g.gridSize), float32(g.gridSize), float32(g.gridSize), clr, true)
	}
	for i, point := range g.snake2.GetPoints() {
		clr = model.ColorSnakeBody2
		if i == 0 {
			clr = model.ColorSnakeHead2
		}
		vector.DrawFilledRect(screen, float32(point.X*g.gridSize), float32(point.Y*g.gridSize), float32(g.gridSize), float32(g.gridSize), clr, true)
	}
}

func (g *game2) drawScore(screen *ebiten.Image) {
	score := fmt.Sprintf("%d", g.snake1.GetScore())
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(g.gridSize), float64(g.gridSize))
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, score, model.Font12, options)

	score = fmt.Sprintf("%d", g.snake2.GetScore())
	options = &text.DrawOptions{}
	options.GeoM.Translate(float64(g.width*g.gridSize-g.gridSize*2), float64(g.gridSize))
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, score, model.Font12, options)
}

func (g *game2) drawGameOver(screen *ebiten.Image) {
	gameOverText := "Game Over"
	textWidth, textHeight := text.Measure(gameOverText, model.Font24, model.Font24.Size)
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(g.width*g.gridSize/2)-float64(textWidth/2), float64(g.height*g.gridSize/2)-float64(textHeight/2)-50)
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, gameOverText, model.Font24, options)

	if g.winner == 0 {
		scoreText := "No winners"
		textWidth, textHeight = text.Measure(scoreText, model.Font12, model.Font12.Size)
		options = &text.DrawOptions{}
		options.GeoM.Translate(float64(g.width*g.gridSize/2)-float64(textWidth/2), float64(g.height*g.gridSize/2)-float64(textHeight/2)-5)
		options.ColorScale.ScaleWithColor(model.ColorText)
		text.Draw(screen, scoreText, model.Font12, options)
	} else {
		scoreText := fmt.Sprintf("Player %d won", g.winner)
		textWidth, textHeight = text.Measure(scoreText, model.Font12, model.Font12.Size)
		options = &text.DrawOptions{}
		options.GeoM.Translate(float64(g.width*g.gridSize/2)-float64(textWidth/2), float64(g.height*g.gridSize/2)-float64(textHeight/2)-5)
		options.ColorScale.ScaleWithColor(model.ColorText)
		text.Draw(screen, scoreText, model.Font12, options)
	}

	pressEnterText := "Press Enter to continue"
	textWidth, textHeight = text.Measure(pressEnterText, model.Font12, model.Font12.Size)
	options = &text.DrawOptions{}
	options.GeoM.Translate(float64(g.width*g.gridSize/2)-float64(textWidth/2), float64(g.height*g.gridSize/2)-float64(textHeight/2)+30)
	options.ColorScale.ScaleWithColor(model.ColorText)
	text.Draw(screen, pressEnterText, model.Font12, options)
}
