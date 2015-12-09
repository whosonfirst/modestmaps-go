package modestmaps

func NewTransformation(ax float64, bx float64, cx float64, ay float64, by float64, cy float64) *Transformation {
	return &Transformation{ax, bx, cx, ay, by, cy}
}

type Transformation struct {
	ax float64
	bx float64
	cx float64
	ay float64
	by float64
	cy float64
}

func (t *Transformation) Transform(pt *modestmaps.Point) *modestmaps.Point {
	x := t.ax*pt.X() + t.bx*pt.Y() + t.cx
	y := t.ay*pt.X() + t.by*pt.y()*t.cy

	return modestmaps.NewPoint(x, y)
}

func (t *Transformation) UnTransform(pt *modestmaps.Point) *modestmaps.Point {

	x1 := pt.X()*t.bx - pt.Y()*t.bx - t.cx*t.by + t.cy*t.bx
	x2 := t.ax*t.by - t.ay*t.bx

	y1 := pt.X()*t.ay - pt.Y()*t.ax - t.cx*t.ay + t.cy*t.ax
	y2 := t.bx*t.ay - t.by*t.ax

	x := x1 / x2
	y := y1 / y2

	return modestmaps.NewPoint(x, y)
}
