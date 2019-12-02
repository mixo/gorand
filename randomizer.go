package gorand

import (
    "time"
    "math/rand"
)

// Returns a random number between min and max numbers
func GetRandBetween(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min) + min
}
