package math

type Point struct {
	x, y float64
}

func (p Point) DistanceToOrigin() float64 {
	return p.x*p.x + p.y*p.y
}
