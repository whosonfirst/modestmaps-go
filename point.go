package modestmaps

import (
	"fmt"
)

func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

type Point struct {
	x float64
	y float64
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) Stringer() string {
	return fmt.Sprintf("%.3f, %.3f", p.X(), p.Y())
}
