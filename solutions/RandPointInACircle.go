package solutions

import (
	"math"
	"math/rand"
)

type Circle struct {
	radius   float64
	x_center float64
	y_center float64
}

func NewCircle(radius float64, x_center float64, y_center float64) Circle {
	return Circle{radius, x_center, y_center}
}

func (c *Circle) RandPoint() []float64 {
	l := c.radius * math.Sqrt(rand.Float64())
	deg := 2 * math.Pi * rand.Float64()
	xRand := c.x_center + l*math.Cos(deg)
	yRand := c.y_center + l*math.Sin(deg)
	return []float64{xRand, yRand}
}
