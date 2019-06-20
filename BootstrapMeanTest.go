package statTest

import (
	"math"
	"math/rand"
	"statTest/stat"
)

// TestResult describes results of statistical test contaning p-value,
// confidence interval and ...
type TestResult struct {
	H0     string
	PValue float64
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

	tOk := 0
	for _, t := range bootstrapT {
		if t >= initT {
			tOk++
		}
	}
	pValue := float64(tOk) / float64(nSim)

	return TestResult{"", pValue}
}

// BootstrapMeanAsync is copy of BootstrapMean but simulations are run in goroutines.
func BootstrapMeanAsync(x, y []float64, nSim int) (TestResult, error) {
	return TestResult{}, nil
}

// TODO: Implementation for single sample mean test.
func BootstrapMeanSingle(x []float64, target float64, nSim int) TestResult {
	// TODO
	return TestResult{}
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
