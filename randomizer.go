package gorand

import (
    "time"
    "math"
    "math/rand"
)

// Returns a random number between min and max numbers
func GetIntRandBetween(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min) + min
}

func GetFloatRandBetween(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	v := min + rand.Float64() * (max - min)
	return math.Round(v * 100) / 100
}
