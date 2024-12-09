package pkg

import (
	"math/rand/v2"
)

func GetRandomFloat(lowerBound, upperBound float64) float64 {
	num := rand.Float64() //nolint
	return lowerBound + (upperBound-lowerBound)*num
}
