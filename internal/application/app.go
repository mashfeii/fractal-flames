package application

import (
	"image"
	"image/png"
	"log/slog"
	"math"
	"math/rand/v2"
	"os"

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

func Init(settings *config.Config) (*App, error) {
	app := &App{
		settings: settings,
	}
	app.generateCoeffs()
	err := app.collectTransitions()
	if err != nil {
		return nil, err
	}

	return app, nil
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

func (a *App) Render() image.Image {
	picture := domain.NewFractalImage(a.settings.Width, a.settings.Height)
	step := 0

	for i := 0; i < a.settings.Samples; i++ {
		newX := pkg.GetRandomFloat(config.XMin, config.XMax)
		newY := pkg.GetRandomFloat(config.YMin, config.YMax)

		for j := -20; j < a.settings.ItNum; j++ {
			chosenTrans := step % len(a.transitions)
			step++

			coeff_matrix := rand.IntN(len(a.coeffs) - 1) //nolint

			affineCoeffs := a.coeffs[coeff_matrix].Affine

			x := affineCoeffs[0]*newX + affineCoeffs[1]*newY + affineCoeffs[2]
			y := affineCoeffs[3]*newX + affineCoeffs[4]*newY + affineCoeffs[5]

			newX, newY = a.transitions[chosenTrans].Convert(x, y)

			if step < 0 {
				continue
			}

			var x1, y1 int

			var xRot, yRot float64

			theta2 := 0.0

			for s := 0; s < a.settings.Symmetry; s++ {
				theta2 += ((2.0 * math.Pi) / float64(a.settings.Symmetry))
				xRot = newX*math.Cos(theta2) - newY*math.Sin(theta2)
				yRot = newX*math.Sin(theta2) + newY*math.Cos(theta2)

				if xRot >= config.XMin && xRot <= config.XMax && yRot >= config.YMin && yRot <= config.YMax {
					x1 = int(float64(a.settings.Width) - ((float64(config.XMax)-xRot)/config.XRange)*float64(a.settings.Width))
					y1 = int(float64(a.settings.Height) - ((float64(config.YMax)-yRot)/config.YRange)*float64(a.settings.Height))
				}

				if picture.Contains(x1, y1) {
					var r, g, b uint8

					pixel := picture.GetPixel(x1, y1)

					if !pixel.Hitted() {
						r = a.coeffs[coeff_matrix].Color[0]
						g = a.coeffs[coeff_matrix].Color[1]
						b = a.coeffs[coeff_matrix].Color[2]
					} else {
						r = (pixel.R + a.coeffs[coeff_matrix].Color[0]) / 2
						g = (pixel.G + a.coeffs[coeff_matrix].Color[1]) / 2
						b = (pixel.B + a.coeffs[coeff_matrix].Color[2]) / 2
					}

					pixel.Hit()
					pixel.SetColor(r, g, b)
				}
			}
		}
	}

	return picture
}

func (a *App) Save(picture image.Image) {
	file, err := os.Create("fractal.png")
	if err != nil {
		slog.Error(err.Error())
	}

	if err := png.Encode(file, picture); err != nil {
		slog.Error(err.Error())
	}

	if err := file.Close(); err != nil {
		slog.Error(err.Error())
	}
}
