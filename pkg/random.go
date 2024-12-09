package pkg

import (
	"math/rand/v2"
)

func GetRandomFloat(lowerBound, upperBound float64) float64 {
	num := rand.Float64()
	return lowerBound + (upperBound-lowerBound)*num
}

// func GetRandomNumber[T int | float64](lower, upper T) T {
// }
