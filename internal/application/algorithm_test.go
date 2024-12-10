package application_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/config"
	"github.com/stretchr/testify/assert"
)

func BenchmarkAlgorithmParallel_5000(b *testing.B) {
	settings, _ := config.NewConfig([6]int{
		1920,
		1080,
		5000,
		20000,
		1,
		12,
	},
		[3]int{
			-1, -1, -1,
		},
		2.2, "", "png")

	err := application.Run(settings)

	assert.NoError(b, err)
}

func BenchmarkAlgorithmSingle_5000(b *testing.B) {
	settings, _ := config.NewConfig([6]int{
		1920,
		1080,
		5000,
		20000,
		1,
		1,
	},
		[3]int{
			-1, -1, -1,
		},
		2.2, "", "png")

	err := application.Run(settings)

	assert.NoError(b, err)
}

func BenchmarkAlgorithmParallel_10000(b *testing.B) {
	settings, _ := config.NewConfig([6]int{
		1920,
		1080,
		10000,
		20000,
		1,
		12,
	},
		[3]int{
			-1, -1, -1,
		},
		2.2, "", "png")

	err := application.Run(settings)

	assert.NoError(b, err)
}

func BenchmarkAlgorithmSingle_10000(b *testing.B) {
	settings, _ := config.NewConfig([6]int{
		1920,
		1080,
		10000,
		20000,
		1,
		1,
	},
		[3]int{
			-1, -1, -1,
		},
		2.2, "", "png")

	err := application.Run(settings)

	assert.NoError(b, err)
}
