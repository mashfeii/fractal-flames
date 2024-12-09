package domain

import (
	"image"
	"image/color"
)

type FractalImage struct {
	width  int
	height int
	pixels []Pixel
	canvas image.Rectangle
}

type Fractal interface {
	image.Image
	GetPixel(x, y int) *Pixel
	Contains(x, y int) bool
}

func NewFractalImage(width, height int) *FractalImage {
	return &FractalImage{
		width:  width,
		height: height,
		pixels: make([]Pixel, width*height),
		canvas: image.Rect(0, 0, width, height),
	}
}

func (f FractalImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (f FractalImage) Bounds() image.Rectangle {
	return f.canvas
}

func (f FractalImage) At(x, y int) color.Color {
	return f.pixels[y*f.width+x]
}

func (f FractalImage) Contains(x, y int) bool {
	return x >= 0 && y >= 0 && x < f.width && y < f.height
}

func (f FractalImage) GetPixel(x, y int) *Pixel {
	return &f.pixels[y*f.width+x]
}
