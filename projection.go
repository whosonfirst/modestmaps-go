package modestmaps

import (
	"math"
)

type IProjection interface {
	RawProject(pt *Point) *Point
	RawUnProject(pt *Point) *Point
	Project(pt *Point) *Point
	UnProject(pt *Point) *Point
	LocationCoordinate(l *Location) *Coordinate
	CoordinateLocation(c *Coordinate) *Location
}

func NewProjectionWithTransformation(zoom int, tranformation *Transformation) *Projection {
	return &Projection{zoom, transformation}
}

func NewProjection(zoom int) *Projection {
	transformation := NewTransformation(1.0, 0.0, 0.0, 0.0, 1.0, 0.0)
	return &Projection{zoom, transformation}
}

type Projection struct {
	IProjection
	zoom           int
	transformation *Transformation
}

func (p *Projection) Project(pt *Point) *Point {
	pt2 := p.RawProject(pt)
	pt3 := p.transformation.Transform(pt2)
	return pt3
}

func (p *Projection) UnProject(pt *Point) *Point {
	pt2 := p.transformation.UnTransform(pt)
	pt3 := p.RawUnProject(pt2)
	return pt3
}

func (p *Projection) LocationCoordinate(l *Location) *Coordinate {

	x := math.PI * l.Lon() / 180.0
	y := math.PI * l.Lat() / 180.0

	pt := NewPoint(x, y)
	pt = p.Project(pt)

	return NewCoordinate(pt.Y(), pt.X(), p.zoom)
}

func (p *Projection) CoordinateLocation(c *Coordinate) *Location {

	c2 := c.ZoomTo(p.zoom)

	pt := NewPoint(c2.column, c2.row)
	pt2 := p.UnProject(pt)

	lat := 180.0 * point.Y() / math.PI
	lon := 180.0 * point.X() / math.PI

	return NewLocation(lat, lon)
}
