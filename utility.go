package main

import (
	// "fmt"
	"math"
)

func DigitLength(num int) (length int) {
	if num == 0 || num == 1 {
		return 1
	} else {
		return int(math.Ceil(math.Log10(float64(num))))
	}
}