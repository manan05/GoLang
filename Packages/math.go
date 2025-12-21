package main

import (
	"math"
	"math/rand"
)

func GetRandom() int {
	return rand.Intn(10)
}

func squareRoot32() float32 {
	return float32(math.Sqrt(2))
}

func squareRoot64() float64 {
	return float64(math.Sqrt(2))
}
