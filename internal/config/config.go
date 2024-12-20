package config

import (
	"strconv"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/errors"
)

const (
	XMax   = 1.77
	XMin   = -1.77
	YMax   = 1.0
	YMin   = -1.0
	XRange = XMax - XMin
	YRange = YMax - YMin
)

type Config struct {
	Width       int
	Height      int
	ItNum       int
	Samples     int
	Symmetry    int
	Threads     int
	NumVectors  int
	Transitions []int
	Colors      [3]int
	Correction  float64
	Format      string
}

func NewConfig(settings [7]int, colors [3]int, corr float64, trans, format string) (*Config, error) {
	for _, val := range settings {
		if val <= 0 {
			return nil, errors.NewErrInvalidIntegerFlag()
		}
	}

	for i := 0; i != len(colors); i++ {
		if colors[i] < -1 || colors[i] > 255 {
			colors[i] = -1
		}
	}

	transitions := make([]int, 0)

	if trans != "" {
		for _, val := range strings.Split(trans, ",") {
			converted, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}

			transitions = append(transitions, converted)
		}
	}

	if format != "jpeg" && format != "png" {
		format = "jpeg"
	}

	return &Config{
		Width:       settings[0],
		Height:      settings[1],
		ItNum:       settings[2],
		Samples:     settings[3],
		Symmetry:    settings[4],
		Threads:     settings[5],
		NumVectors:  settings[6],
		Colors:      colors,
		Transitions: transitions,
		Correction:  corr,
		Format:      format,
	}, nil
}
