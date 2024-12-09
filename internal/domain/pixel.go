package domain

type FractalPixel interface {
	Hitted() bool
	Hit()
	SetColor(r, g, b uint8)
}

type Pixel struct {
	R, G, B  uint8
	HitCount uint
}

func (p Pixel) RGBA() (r, g, b, a uint32) {
	r = uint32(p.R)
	r |= r << 8
	g = uint32(p.G)
	g |= g << 8
	b = uint32(p.B)
	b |= b << 8
	a = 255
	a |= a << 8

	return
}

func (p Pixel) Hitted() bool {
	return p.HitCount != 0
}

func (p *Pixel) Hit() {
	p.HitCount++
}

func (p *Pixel) SetColor(r, g, b uint8) {
	p.R = r
	p.G = g
	p.B = b
}
