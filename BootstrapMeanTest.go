package statTest

import (
	"math"
	"math/rand"
	"statTest/stat"
)

// BootstrapMeanH0 is base hypothesis for BootstrapMean test.
const BootstrapMeanH0 = `H0: Mean(x) == Mean(y), without assumming distribution of x or y.`

// BootstrapMeanSingleH0 is base hypothesis for BootstrapMeanSingle test.
const BootstrapMeanSingleH0 = `H0: Mean(x) == u, without assumming distribution of x or its variance`

// TestResult describes results of statistical test contaning p-value,
// confidence interval and ...
type TestResult struct {
	H0     string
	PValue float64
	CI     ConfInterval
}

// ConfInterval struct represents confidence interval of statistical tests.
type ConfInterval struct {
	LowerBound float64
	UpperBound float64
}

// BootstrapMean is a nonparametric statistical test for testing equality
// of means of two independent data samples. In this implementation data samples
// are expressed as slices of float64. Argument nSim describes number of
// simulations (nSim = 1000 is usually default value).
// Implementation is based on algorithm proposed by Efron and Tibshirani in 1993
// in paper "An Introduction to the Bootstrap".
// TODO: What about the seed?
func BootstrapMean(x, y []float64, nSim int) TestResult {
	initT := bootstrapTStat(x, y)
	joint := append(x, y...)
	jointMean := stat.Mean(joint)
	bootstrapT := make([]float64, nSim)
	xMean := stat.Mean(x)
	yMean := stat.Mean(y)
	xPrime := addConst(x, -1.0*xMean+jointMean)
	yPrime := addConst(y, -1.0*yMean+jointMean)

	for i := 0; i < nSim; i++ {
		newX := SampleRepl(xPrime)
		newY := SampleRepl(yPrime)
		bootstrapT[i] = bootstrapTStat(newX, newY)
	}

	tOk := countOver(bootstrapT, initT)
	pValue := float64(tOk) / float64(nSim)

	return TestResult{BootstrapMeanH0, pValue, ConfInterval{math.NaN(),
		math.NaN()}}
}

// BootstrapMeanAsync is copy of BootstrapMean but simulations are run in goroutines.
func BootstrapMeanAsync(x, y []float64, nSim int) (TestResult, error) {
	return TestResult{}, nil
}

// BootstrapMeanSingle performes single-sample bootstrap mean test for testing
// that mean of given sample is equal to given number u0. Parameter alpha set
// power of confidence interval and nSim states number of simulation in bootstrap
// performed during the test. Test does not assume anything about sample x, its
// distributions, variance etc.
// Implementation is based on algorithm proposed by Efron and Tibshirani in 1993
// in paper "An Introduction to the Bootstrap".
// TODO: What about the seed?
func BootstrapMeanSingle(x []float64, u0, alpha float64, nSim int) TestResult {
	bootstrapT := make([]float64, nSim)
	xMean := stat.Mean(x)
	xPrime := addConst(x, -1.0*xMean+u0)
	initT := bootstrapTStatSingle(x, u0)

	for i := 0; i < nSim; i++ {
		newX := SampleRepl(xPrime)
		bootstrapT[i] = bootstrapTStatSingle(newX, u0)
	}

	tOk := countOver(bootstrapT, initT)
	pValue := float64(tOk) / float64(nSim)
	ci := bootstrapMeanSingleXCI(x, bootstrapT, alpha)

	return TestResult{BootstrapMeanSingleH0, pValue, ci}
}

// Function bootstrapMeanSingleXCI calculates confidence interval for single
// bootstrap mean test. Confidence interval is of the form:
//
// 	(xMean - t_{1 - alpha} * sd, xMean - t_{alpha} * sd)
//
// where:
// 	xMean is mean of x
// 	t_{a} is a-quantile of vector ts
// 	sd is standard deviation.
func bootstrapMeanSingleXCI(x []float64, bootstrapT []float64,
	alpha float64) ConfInterval {
	if alpha < 0.0 || alpha > 1.0 || math.IsNaN(alpha) || math.IsInf(alpha, 0) {
		return ConfInterval{math.Inf(-1), math.Inf(1)}
	}

	xMean := stat.Mean(x)
	xSd := math.Sqrt(stat.Variance(x)) / math.Sqrt(float64(len(x)))
	lower := xMean - stat.Quantile(bootstrapT, 1.0-alpha)*xSd
	upper := xMean - stat.Quantile(bootstrapT, alpha)*xSd

	return ConfInterval{lower, upper}
}

// Function calcTStat calculates statistic t:
// t = \frac{mean(x) - mean(y)}{\sqrt{\var{x} / n + \var{y} / m}}
func bootstrapTStat(x, y []float64) float64 {
	if len(x) == 0 || len(y) == 0 {
		return math.NaN()
	}
	xMean := stat.Mean(x)
	yMean := stat.Mean(y)
	xVar := stat.Variance(x)
	yVar := stat.Variance(y)
	xl := 1.0 / float64(len(x))
	yl := 1.0 / float64(len(y))

	return (xMean - yMean) / math.Sqrt(xVar*xl+yVar*yl)
}

// Function bootstrapTStatSingle calculates bootstrap t-statistic for single
// sample mean bootstrap test:
// t = \frac{mean(x) - u0}{\sqrt{\var{x} * \frac{1}{n}}}.
func bootstrapTStatSingle(x []float64, u0 float64) float64 {
	if len(x) == 0 {
		return math.NaN()
	}
	xMean := stat.Mean(x)
	xVar := stat.Variance(x)
	xl := 1.0 / float64(len(x))
	return (xMean - u0) / math.Sqrt(xVar*xl)
}

// SampleRepl samples with replacements from given slice of float64.
// TODO: What about a seed?
func SampleRepl(x []float64) []float64 {
	xLen := len(x)
	sampledX := make([]float64, xLen)

	for i := 0; i < xLen; i++ {
		id := rand.Intn(xLen)
		sampledX[i] = x[id]
	}
	return sampledX
}

// addConst adds to every element of x value of given constans cnst.
func addConst(x []float64, cnst float64) []float64 {
	newX := make([]float64, len(x))
	for i, e := range x {
		newX[i] = e + cnst
	}
	return newX
}

func countOver(x []float64, t float64) int {
	tOk := 0
	for _, e := range x {
		if e >= t {
			tOk++
		}
	}
	return tOk
}
