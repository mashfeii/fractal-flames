package application

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/errors"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/transformations"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

type App struct {
	settings    *config.Config
	coeffs      []domain.Coeff
	transitions []transformations.Transformation
}

func (a *App) generateCoeffs() {
	for i := 0; i < a.settings.N; i++ {
		newCoeff := domain.Coeff{}

		for i := 0; i != 6; i++ {
			newCoeff.Affine[i] = pkg.GetRandomFloat(-1.5, 1.5)
		}

		for i := 0; i != 3; i++ {
			newCoeff.Color[i] = uint8(pkg.GetRandomFloat(0, 255))
		}

		a.coeffs = append(a.coeffs, newCoeff)
	}
}

func (a *App) collectTransitions() error {
	possibleTransitions := []transformations.Transformation{
		&transformations.Disc{},
		&transformations.Heart{},
		&transformations.Handkerchief{},
		&transformations.Polar{},
		&transformations.Spiral{},
		&transformations.Spherical{},
		&transformations.Swirl{},
		&transformations.Hyperbolic{},
	}

	if len(a.settings.Transitions) == 0 {
		a.transitions = possibleTransitions
		return nil
	}

	for _, val := range a.settings.Transitions {
		if val < 0 || val >= len(possibleTransitions) {
			continue
		}

		a.transitions = append(a.transitions, possibleTransitions[val])
	}

	if len(a.transitions) == 0 {
		return errors.NewErrEmptyTransitions(a.settings.Transitions)
	}

	return nil
}

func (a *App) iterate(coeffs, trans int, x, y float64) (newX, newY float64) {
	affineCoeffs := a.coeffs[coeffs].Affine

	// NOTE: Apply vector coefficient
	Xcoeff := affineCoeffs[0]*x + affineCoeffs[1]*y + affineCoeffs[2]
	Ycoeff := affineCoeffs[3]*x + affineCoeffs[4]*y + affineCoeffs[5]

	// NOTE: Apply transformation
	return a.transitions[trans].Convert(Xcoeff, Ycoeff)
}

func (a *App) drawPixel(pic *domain.FractalImage, coeffs, x1, y1 int) {
	var r, g, b uint8

	pixel := pic.GetPixel(x1, y1)

	if !pixel.Hitted() {
		r = a.coeffs[coeffs].Color[0]
		g = a.coeffs[coeffs].Color[1]
		b = a.coeffs[coeffs].Color[2]
	} else {
		r = (pixel.R + a.coeffs[coeffs].Color[0]) / 2
		g = (pixel.G + a.coeffs[coeffs].Color[1]) / 2
		b = (pixel.B + a.coeffs[coeffs].Color[2]) / 2
	}

	pixel.Hit()
	pixel.SetColor(r, g, b)
}

func (a *App) render() domain.Fractal {
	picture := domain.NewFractalImage(a.settings.Width, a.settings.Height)
	step := 0

	for i := 0; i < a.settings.Samples; i++ {
		// NOTE: Take new point from the canvas
		newX := pkg.GetRandomFloat(config.XMin, config.XMax)
		newY := pkg.GetRandomFloat(config.YMin, config.YMax)

		for j := -20; j < a.settings.ItNum; j++ {
			// NOTE: Take the next transformation
			chosenTrans := step % len(a.transitions)
			step++

			coeff_matrix := rand.IntN(len(a.coeffs) - 1) //nolint

			newX, newY = a.iterate(coeff_matrix, chosenTrans, newX, newY)

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
					a.drawPixel(picture, coeff_matrix, x1, y1)
				}
			}
		}
	}

	return picture
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

func (a *App) save(picture image.Image) error {
	if err := os.MkdirAll(filepath.Join(".", "results"), os.ModePerm); err != nil {
		return err
	}

	entities, err := os.ReadDir("results/")
	if err != nil {
		return err
	}

	counter := 0

	for _, entity := range entities {
		if strings.Contains(entity.Name(), "fractal") {
			counter++
		}
	}

	path := "results/fractal." + a.settings.Format
	if counter > 0 {
		path = fmt.Sprintf("results/fractal-%d.%s", counter, a.settings.Format)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	switch a.settings.Format {
	case "jpeg":
		if err := jpeg.Encode(file, picture, nil); err != nil {
			return err
		}
	default:
		if err := png.Encode(file, picture); err != nil {
			return err
		}
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}
