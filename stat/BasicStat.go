// Package stat contains low-level statistical functions.
package stat

import (
	"math"
)

// Sum returns sum of elements from float64 slice.
func Sum(x []float64) float64 {
	s := 0.0
	for _, e := range x {
		s += e
	}
	return s
}

// Len is wrapper for len which return float64.
func Len(x []float64) float64 {
	return float64(len(x))
}

// Mean calculates mean of given float64 vector. In case of input of
// length 0 it returns 'not-a-number' float64.
func Mean(x []float64) float64 {
	if len(x) == 0 {
		return math.NaN()
	}
	return Sum(x) / Len(x)
}

// Variance calculates mathematical variance - E((X - mean)^2).
func Variance(x []float64) float64 {
	if len(x) <= 1 {
		return math.NaN()
	}

	m := Mean(x)
	v := 0.0

	for _, e := range x {
		v += math.Pow(e-m, 2.0)
	}
	return v / (Len(x) - 1)
}

// ToPower function takes to power value all elements of slice x.
func ToPower(x []float64, power float64) []float64 {
	xToPow := make([]float64, len(x))
	for i, e := range x {
		xToPow[i] = math.Pow(e, power)
	}
	return xToPow
}
