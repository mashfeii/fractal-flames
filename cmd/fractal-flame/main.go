package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/config"
)

func main() {
	width := flag.Int("w", 1920, "Width for the future picture")
	height := flag.Int("h", 1080, "Height for the future picture")
	samples := flag.Int("s", 20000, "Number of samples to generate picture")
	iterations := flag.Int("i", 1000, "Number of iterations per sample")
	coeffVectors := flag.Int("n", 24, "Number of random vectors of coefficients")
	symmetry := flag.Int("sym", 1, "Number of symmetry axis")
	transitions := flag.String("t", "0,2,3", "List of transitions separated by comma")

	flag.Parse()

	settings, err := config.NewConfig(
		[]int{
			*width,
			*height,
			*iterations,
			*samples,
			*coeffVectors,
			*symmetry,
		}, *transitions)
	if err != nil {
		slog.Error(err.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}

	driver, err := application.Init(settings)
	if err != nil {
		slog.Error(err.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}

	generatedImage := driver.Render()

	driver.Save(generatedImage)
}
