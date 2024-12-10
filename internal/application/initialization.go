package application

import (
	"sync"

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
	mutex       sync.Mutex
}

func (a *App) generateCoeffs() {
	for i := 0; i < len(a.transitions); i++ {
		newCoeff := domain.Coeff{}

		for i := 0; i != 6; i++ {
			newCoeff.Affine[i] = pkg.GetRandomFloat(-1.5, 1.5)
		}

		for i := 0; i != 3; i++ {
			newCoeff.Color[i] = uint8(pkg.GetRandomFloat(64, 255))
		}

		a.coeffs = append(a.coeffs, newCoeff)
	}
}

func (a *App) collectTransitions() error {
	possibleTransitions := []transformations.Transformation{
		&transformations.Sinusoidal{},
		&transformations.Spherical{},
		&transformations.Swirl{},
		&transformations.Horseshoe{},
		&transformations.Polar{},
		&transformations.Handkerchief{},
		&transformations.Heart{},
		&transformations.Disc{},
		&transformations.Spiral{},
		&transformations.Hyperbolic{},
		&transformations.Ex{},
		&transformations.Julia{},
		&transformations.Fisheye{},
		&transformations.Eyefish{},
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
