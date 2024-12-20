package application

import (
	"math"
	"math/rand/v2"
	"sync"

	"github.com/es-debug/backend-academy-2024-go-template/internal/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

func (a *App) iterate(choice, trans int, x, y float64) (newX, newY float64) {
	affineCoeffs := a.coeffs[choice].Affine

	// NOTE: Apply vector coefficient
	Xcoeff := affineCoeffs[0]*x + affineCoeffs[1]*y + affineCoeffs[2]
	Ycoeff := affineCoeffs[3]*x + affineCoeffs[4]*y + affineCoeffs[5]

	// NOTE: Apply transformation
	return a.transitions[trans].Convert(Xcoeff, Ycoeff)
}

func (a *App) drawPixel(pic *domain.FractalImage, coeffs, x1, y1 int) {
	pixel := pic.GetPixel(x1, y1)

	r := pixel.R/2 + a.coeffs[coeffs].Color[0]/2
	g := pixel.G/2 + a.coeffs[coeffs].Color[1]/2
	b := pixel.B/2 + a.coeffs[coeffs].Color[2]/2

	pixel.Hit()
	pixel.SetColor(r, g, b)
}

func (a *App) render() domain.Fractal {
	picture := domain.NewFractalImage(a.settings.Width, a.settings.Height)
	routinesGroup := sync.WaitGroup{}

	for i := 0; i < a.settings.Threads; i++ {
		routinesGroup.Add(1)

		go func() {
			a.renderStep(picture)
			routinesGroup.Done()
		}()
	}

	routinesGroup.Wait()

	return picture
}

func (a *App) renderStep(picture *domain.FractalImage) {
	for i := 0; i < a.settings.Samples; i++ {
		// NOTE: Take new point from the canvas
		newX := pkg.GetRandomFloat(config.XMin, config.XMax)
		newY := pkg.GetRandomFloat(config.YMin, config.YMax)

		for j := -20; j < a.settings.ItNum; j++ {
			// NOTE: Take the next transformation
			trans := rand.IntN(len(a.transitions)) //nolint
			choice := rand.IntN(len(a.coeffs))

			newX, newY = a.iterate(choice, trans, newX, newY)

			// NOTE: Skip first iterations to find the center
			if j < 0 {
				continue
			}

			theta2 := 0.0

			for s := 0; s < a.settings.Symmetry; s++ {
				var x1, y1 int

				// NOTE: Apply symmetry (rotation)
				theta2 += ((2.0 * math.Pi) / float64(a.settings.Symmetry))
				xRot := newX*math.Cos(theta2) - newY*math.Sin(theta2)
				yRot := newX*math.Sin(theta2) + newY*math.Cos(theta2)

				if xRot >= config.XMin && xRot <= config.XMax && yRot >= config.YMin && yRot <= config.YMax {
					x1 = int(float64(a.settings.Width) - ((float64(config.XMax)-xRot)/config.XRange)*float64(a.settings.Width))
					y1 = int(float64(a.settings.Height) - ((float64(config.YMax)-yRot)/config.YRange)*float64(a.settings.Height))
				}

				// NOTE: "Plot picture" -> Set pixel color based on hit count
				if picture.Contains(x1, y1) {
					a.mutex.Lock()
					a.drawPixel(picture, choice, x1, y1)
					a.mutex.Unlock()
				}
			}
		}
	}
}

func (a *App) correction(picture domain.Fractal) {
	maxNormalized := 0.0

	// NOTE: Apply logarithmic correction
	for x := 0; x != a.settings.Width; x++ {
		for y := 0; y != a.settings.Height; y++ {
			pixel := picture.GetPixel(x, y)

			if !pixel.Hitted() {
				continue
			}

			maxNormalized = math.Max(maxNormalized, pixel.Normalize())
		}
	}

	for x := 0; x != a.settings.Width; x++ {
		for y := 0; y != a.settings.Height; y++ {
			picture.GetPixel(x, y).Correction(maxNormalized, a.settings.Correction)
		}
	}
}
