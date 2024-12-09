package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/config"
)

func main() {
	width := flag.Int("w", 1024, "Width for the future picture")
	height := flag.Int("h", 768, "Height for the future picture")
	samples := flag.Int("s", 20000, "Number of samples to generate picture")
	iterations := flag.Int("i", 1000, "Number of iterations per sample")
	coeffVectors := flag.Int("n", 20, "Number of random vectors of coefficients")
	symmetry := flag.Int("sym", 1, "Number of symmetry axis")
	transitions := flag.String("t", "1,4,6", "List of transitions separated by comma")
	gammaCorrection := flag.Float64("g", 2.2, "Gamma correction coefficient (0 - disable)")
	format := flag.String("f", "png", "Image format")

	flag.Parse()

	settings, err := config.NewConfig(
		[]int{
			*width,
			*height,
			*iterations,
			*samples,
			*coeffVectors,
			*symmetry,
		}, *gammaCorrection, *transitions, *format)
	if err != nil {
		slog.Error(err.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := application.Run(settings); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
