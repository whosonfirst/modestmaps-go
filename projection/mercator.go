package modestmaps

import (
       "math"
)

func NewMercatorProjection (zoom int) *MercatorProjection {
     return &MercatorProjection(zoom)
}

type MercatorProjection struct {
     Projection
}

func (m *MercatorProjection) RawProject (pt *Point) *Point {

     x := pt.X()
     y := math.Log(math.Tan(0.25 * math.PI + 0.5 * pt.Y()))

     return NewPoint(x, y)
}

func (m *MercatorProjection) RawUnProject (pt *Point) *Point {

     x := pt.X()
     y := 2.0 * math.Atan(math.Pow(math.e, pt.Y())) - 0.5 * math.PI

     return NewPoint(x, y)
}
