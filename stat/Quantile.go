package stat

import (
	"math"
	"sort"
)

// Quantile calculates quantile of rank prob based on given sample x. This
// implementation is consistent with default implemenatation in function
// "quantile" in R programming language. Algorithm was described in
// "Sample Quantiles in Statistical Packages" - Hyndman and Fan (1996) as
// "Definition 7". Quantile estimator can be expressed as follows:
//
// 	Q(p) = (1 - g(p))*X_{(j)} + g(p)*X_{(j + 1)}
//
// where
// 	X_{(j)} denotes j-th order statistic of X = (X_1, X_2, ..., X_n),
// 	j = floor{p(n - 1) + 1},
// 	g(p) = p(n - 1) + 1 - floor{p(n - 1) + 1},
// 	and 0 <= p <= 1.
func Quantile(x []float64, prob float64) float64 {
	if prob < 0.0 || prob > 1.0 {
		return math.NaN()
	}
	newX := sort.Float64Slice(x)
	newX.Sort()

	if prob == 1.0 {
		return newX[len(newX)-1]
	}

	j := qj(x, prob)
	gpVal := gp(x, prob)
	quantile := (1.0-gpVal)*newX[j] + gpVal*newX[j+1]

	return quantile
}

// Quantiles does exactly the same as Quantile but calculates possible many
// quantilies for each given probability in argument probs. The same result
// can be obtained by calling Quantile several times. Difference is that in
// this implementation x slice is sorted only one time, thus for large x it
// can be helpful in term of performance.
func Quantiles(x, probs []float64) []float64 {
	quantiles := make([]float64, len(probs))
	newX := sort.Float64Slice(x)
	newX.Sort()

	for i, prob := range probs {
		if prob < 0.0 || prob > 1.0 {
			quantiles[i] = math.NaN()
			continue
		}
		if prob == 1.0 {
			quantiles[i] = newX[len(newX)-1]
			continue
		}

		j := qj(x, prob)
		gpVal := gp(x, prob)
		quantiles[i] = (1.0-gpVal)*newX[j] + gpVal*newX[j+1]
	}

	return quantiles
}

// Function qj calculates "j" index in order statistics for Quantile function.
func qj(x []float64, prob float64) int {
	xl := float64(len(x))
	stat := prob*(xl-1.0) + 1.0
	return int(math.Floor(stat)) - 1 // -1 because go is 0-index
}

// Function gp calculates value of g(p) function for Quantile function.
func gp(x []float64, prob float64) float64 {
	xl := float64(len(x))
	stat := prob*(xl-1.0) + 1.0
	gpValue := stat - math.Floor(stat)
	return gpValue
}
