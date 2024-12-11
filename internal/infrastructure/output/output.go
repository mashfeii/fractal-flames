package output

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func Save(format string, picture image.Image) error {
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

	path := "results/fractal." + format
	if counter > 0 {
		path = fmt.Sprintf("results/fractal-%d.%s", counter, format)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	switch format {
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
