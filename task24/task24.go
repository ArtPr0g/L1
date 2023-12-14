package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

func (p *Point) DistanceTo(q *Point) float64 {
	dx := p.X - q.X
	dy := p.Y - q.Y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	distance := p1.DistanceTo(p2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
