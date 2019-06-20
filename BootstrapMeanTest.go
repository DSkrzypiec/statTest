package statTest

// TestResult describes results of statistical test contaning p-value,
// confidence interval and ...
type TestResult struct {
}

// BootstrapMeanTest is a nonparametric statistical test for testing equality
// of means of two independent data samples. In this implementation data samples
// are expressed as slices of float64. Argument nSim describes number of
// simulations (nSim = 1000 is usually default value).
// Implementation is based on algorithm proposed by Efron and Tibshirani in 1993
// in paper "An Introduction to the Bootstrap".
func BootstrapMeanTest(x, y []float64, nSim int) TestResult {
	return TestResult{}
}

// Function calcTStat calculates statistic t:
// t = \frac{mean(x) - mean(y)}{\sqrt{\var{x} / n + \var{y} / m}}
func calcTStat(x, y []float64) float64 {
	return 0.0
}
