package transformations

import "math"

type Transformation interface {
	Convert(x, y float64) (newX, newY float64)
}

type Sinusoidal struct{}

func (s *Sinusoidal) Convert(x, y float64) (newX, newY float64) {
	return math.Sin(x), math.Sin(y)
}

type Spherical struct{}

func (s *Spherical) Convert(x, y float64) (newX, newY float64) {
	r := 1.0 / (x*x + y*y)
	return r * x, r * y
}

type Swirl struct{}

func (s *Swirl) Convert(x, y float64) (newX, newY float64) {
	r := x*x + y*y
	return x*math.Sin(r) - y*math.Cos(r), x*math.Cos(r) + y*math.Sin(r)
}

type Horseshoe struct{}

func (h *Horseshoe) Convert(x, y float64) (newX, newY float64) {
	r := 1.0 / math.Sqrt(x*x+y*y)
	return (x - y) * (x + y) * r, 2 * x * y * r
}

type Polar struct{}

func (p *Polar) Convert(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(x, y)

	return theta / math.Pi, r - 1
}

type Handkerchief struct{}

func (h *Handkerchief) Convert(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(x, y)

	return r * math.Sin(theta+r), r * math.Cos(theta-r)
}

type Heart struct{}

func (h *Heart) Convert(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(x, y)

	return r * math.Sin(theta*r), r * (-math.Cos(theta * r))
}

type Disc struct{}

func (d *Disc) Convert(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(x, y)
	mult := theta / math.Pi

	return mult * math.Sin(math.Pi*r), mult * math.Cos(math.Pi*r)
}

type Spiral struct{}

func (s *Spiral) Convert(x, y float64) (newX, newY float64) {
	r := 1.0 / math.Sqrt(x*x+y*y)
	theta := math.Atan2(x, y)

	return r * (math.Cos(theta) + math.Sin(r)), r * (math.Sin(theta) - math.Cos(r))
}

type Hyperbolic struct{}

func (h *Hyperbolic) Convert(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(x, y)

	return math.Sin(theta) / r, r * math.Cos(theta)
}
