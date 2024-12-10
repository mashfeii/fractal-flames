package domain

import "math"

type FractalPixel interface {
	Hitted() bool
	Hit()
	SetColor(r, g, b, a uint8)
}

type Pixel struct {
	R, G, B    uint8
	hitCount   uint
	normalized float64
}

func (p Pixel) RGBA() (r, g, b, a uint32) {
	r = uint32(p.R)
	r |= r << 8
	g = uint32(p.G)
	g |= g << 8
	b = uint32(p.B)
	b |= b << 8
	a = uint32(255)
	a |= a << 8

	return
}

func (p Pixel) Hitted() bool {
	return p.hitCount != 0
}

func (p *Pixel) Hit() {
	p.hitCount++
}

func (p *Pixel) SetColor(r, g, b uint8) {
	p.R = r
	p.G = g
	p.B = b
}

func (p *Pixel) Normalize() float64 {
	p.normalized = math.Log10(float64(p.hitCount))

	return p.normalized
}

func (p *Pixel) Correction(coeff, gamma float64) {
	p.normalized /= coeff
	p.R = uint8(float64(p.R) * math.Pow(p.normalized, 1.0/gamma))
	p.G = uint8(float64(p.G) * math.Pow(p.normalized, 1.0/gamma))
	p.B = uint8(float64(p.B) * math.Pow(p.normalized, 1.0/gamma))
}
