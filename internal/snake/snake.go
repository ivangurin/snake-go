package snake

import "snake-go/internal/model"

func NewSnake(width, height int, direction *model.Point) *Snake {
	return &Snake{
		width:     width,
		height:    height,
		points:    []model.Point{{X: width / 2, Y: height / 2}},
		direction: direction,
	}
}

func (s *Snake) Move() {
	s.points = append([]model.Point{s.getNewHead()}, s.points[:len(s.points)-1]...)
}

func (s *Snake) SetDirection(direction *model.Point) {
	if direction.X == -s.direction.X && direction.Y == -s.direction.Y {
		return
	}
	s.direction = direction
}

func (s *Snake) GetHeadPoint() *model.Point {
	return &s.points[0]
}

func (s *Snake) GetPoints() []model.Point {
	return s.points
}

func (s *Snake) GrowUp() {
	s.points = append([]model.Point{s.getNewHead()}, s.points...)
}

func (s *Snake) getNewHead() model.Point {
	newHead := model.Point{
		X: s.points[0].X + s.direction.X,
		Y: s.points[0].Y + s.direction.Y,
	}

	if newHead.X < 0 {
		newHead.X = s.width - 1
	} else if newHead.X >= s.width {
		newHead.X = 0
	}

	if newHead.Y < 0 {
		newHead.Y = s.height - 1
	} else if newHead.Y >= s.height {
		newHead.Y = 0
	}

	return newHead
}

func (s *Snake) GetScore() int {
	return len(s.points) - 1
}
