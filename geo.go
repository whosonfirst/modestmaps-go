package modestmaps

func DeriveTransformation(a1x float64, a1y float64, a2x float64, a2y float64, b1x float64, b1y float64, b2x float64, b2y float64, c1x float64, c1y float64, c2x float64, c2y float64) *modestmaps.Transformation {

	ax, bx, cx := LinearSolution(a1x, a1y, a2x, b1x, b1y, b2x, c1x, c1y, c2x)
	ay, by, cy := LinearSolution(a1x, a1y, a2y, b1x, b1y, b2y, c1x, c1y, c2y)

	return modestmaps.NewTransformation(ax, bx, cx, ay, by, cy)
}

func LinearSolution(r1 float64, s1 float64, t1 float64, r2 float64, s2 float64, t2 float64, r3 float64, s3 float64, t3 float64) (a float64, b float64, c float64) {

	a = (((t2 - t3) * (s1 - s2)) - ((t1 - t2) * (s2 - s3))) / (((r2 - r3) * (s1 - s2)) - ((r1 - r2) * (s2 - s3)))
	b = (((t2 - t3) * (r1 - r2)) - ((t1 - t2) * (r2 - r3))) / (((s2 - s3) * (r1 - r2)) - ((s1 - s2) * (r2 - r3)))
	c = t1 - (r1 * a) - (s1 * b)

	return a, b, c
}
