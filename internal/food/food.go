package food

import "snake-go/internal/model"

func NewFood(x, y int) *Food {
	return &Food{
		point: &model.Point{
			X: x,
			Y: y,
		},
	}
}

func (f *Food) GetPoint() *model.Point {
	return f.point
}
