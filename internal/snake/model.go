package snake

import "snake-go/internal/model"

type Snake struct {
	width     int
	height    int
	points    []model.Point
	direction *model.Point
}
