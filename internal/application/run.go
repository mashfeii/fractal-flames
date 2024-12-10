package application

import "github.com/es-debug/backend-academy-2024-go-template/internal/config"

func Run(settings *config.Config) error {
	app := App{
		settings: settings,
	}

	if err := app.collectTransitions(); err != nil {
		return err
	}

	app.generateCoeffs()

	generatedImage := app.render()

	if settings.Correction != 0.0 {
		app.correction(generatedImage)
	}

	if err := app.save(generatedImage); err != nil {
		return err
	}

	return nil
}
