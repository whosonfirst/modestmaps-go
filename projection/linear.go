package modestmaps

func NewLinearProjection (zoom int) *LinearProjection {
     return &LinearProjection(zoom)
}

type LinearProjection struct {
     Projection
}

func (l *LinearProjection) RawProject (pt *Point) *Point {
     return NewPoint(pt.X(), pt.y())
}

func (l *LinearProjection) RawUnProject (pt *Point) *Point {
     return NewPoint(pt.X(), pt.y())
}
