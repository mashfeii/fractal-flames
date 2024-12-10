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
	// TODO: Deal with single transformation
	/* coeffVectors := flag.Int("n", 20, "Number of random vectors of coefficients") */ //nolint
	symmetry := flag.Int("sym", 1, "Number of symmetry axis")
	threads := flag.Int("t", 12, "Number of threads")
	transitions := flag.String("tr", "", "List of transitions separated by comma (\"1,2,3\")")
	gammaCorrection := flag.Float64("g", 2.2, "Gamma correction coefficient (0 - disable)")
	format := flag.String("f", "png", "Image format")
	r := flag.Int("re", -1, "Static red channel value")
	g := flag.Int("gr", -1, "Static green channel value")
	b := flag.Int("bl", -1, "Static blue channel value")

	flag.Parse()

	settings, err := config.NewConfig(
		[6]int{
			*width,
			*height,
			*iterations,
			*samples,
			*symmetry,
			*threads,
		},
		[3]int{
			*r, *g, *b,
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
