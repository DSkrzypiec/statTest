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
	newX := sort.Float64Slice(x)
	newX.Sort()
	j := qj(x, prob)
	gpVal := gp(x, prob)

	if j > len(x)-2 {
		return math.NaN()
	}

	quantile := (1.0-gpVal)*newX[j] + gpVal*newX[j+1]
	return quantile
}

// Function qj calculates "j" index in order statistics for Quantile function.
func qj(x []float64, prob float64) int {
	xl := float64(len(x))
	stat := prob*(xl-1.0) + 1.0
	return int(math.Floor(stat))
}

// Function gp calculates value of g(p) function for Quantile function.
func gp(x []float64, prob float64) float64 {
	xl := float64(len(x))
	stat := prob*(xl-1.0) + 1.0
	gpValue := stat - math.Floor(stat)
	return gpValue
}
