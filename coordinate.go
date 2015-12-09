package modestmaps

import (
	"fmt"
	"math"
)

func NewCoordinate(row int, column int, zoom int) *Coordinate {
	return &Coordinate{row, column, zoom}
}

type Coordinate struct {
	row    int
	column int
	zoom   int
}

func (c *Coordinate) Stringer() string {
	return fmt.Sprintf("%.3f, %.3f @%.3f", c.row, c.column, c.zoom)
}

// TO DO __eq__
// TO DO __cmp__
// TO DO __hash__

func (c *Coordinate) Copy() *Coordinate {
	return NewCoordinate(c.row, c.column, c.zoom)
}

func (c *Coordinate) Container() *Coordinate {

	_row := float64(c.row)
	_row = math.Floor(_row)

	_column := float64(c.column)
	_column = math.Floor(_column)

	row := int(_row)
	column := int(_column)
	zoom := c.zoom

	return NewCoordinate(row, column, zoom)
}

func (c *Coordinate) ZoomTo(distance int) *Coordinate {

	_distance := float64(distance)
	_zoom := float64(c.zoom)

	_row := math.Pow(2.0, _distance-_zoom)
	_column := math.Pow(2, _distance-_zoom)

	row := c.row * int(_row)
	column := c.column * int(_column)
	zoom := c.zoom

	return NewCoordinate(row, column, zoom)
}

func (c *Coordinate) ZoomBy(distance int) *Coordinate {

	_distance := float64(distance)

	_row := math.Pow(2, _distance)
	_column := math.Pow(2, _distance)

	row := c.row * int(_row)
	column := c.column * int(_column)
	zoom := c.zoom

	return NewCoordinate(row, column, zoom)
}

func (c *Coordinate) Up(distance int) *Coordinate {
	row := c.row - distance
	column := c.column
	zoom := c.zoom

	return NewCoordinate(row, column, zoom)
}

func (c *Coordinate) Right(distance int) *Coordinate {
	row := c.row
	column := c.column + distance
	zoom := c.zoom

	return NewCoordinate(row, column, zoom)
}

func (c *Coordinate) Down(distance int) *Coordinate {
	row := c.row + distance
	column := c.column
	zoom := c.zoom

	return NewCoordinate(row, column, zoom)
}

func (c *Coordinate) Left(distance int) *Coordinate {
	row := c.row
	column := c.column - distance
	zoom := c.zoom

	return NewCoordinate(row, column, zoom)
}
