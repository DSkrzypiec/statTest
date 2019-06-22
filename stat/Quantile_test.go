package stat

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestUnitQuantileBasic(t *testing.T) {
	x, rResults := quantileResultsFromR1()
	x2, rResults2 := quantileResultsFromR2()

	for p, rQuant := range rResults {
		goQuant := Quantile(x, p)

		if !closeEnough(rQuant, goQuant, 0.001) {
			t.Errorf("Incorrect quantile for p = %f. Expected: %f, got: %f \n",
				p, rQuant, goQuant)
		}
	}

	for p, rQuant := range rResults2 {
		goQuant := Quantile(x2, p)

		if !closeEnough(rQuant, goQuant, 0.001) {
			t.Errorf("Incorrect quantile for p = %f. Expected: %f, got: %f \n",
				p, rQuant, goQuant)
		}
	}
}

func TestUnitQuantileEdge(t *testing.T) {
	wierdX := []float64{math.NaN(), math.Inf(1), 1.0, 5.431}
	nanRes := Quantile(wierdX, 0.12)
	if !math.IsNaN(nanRes) {
		t.Errorf("Expected NaN, got: %f", nanRes)
	}

	infRes := Quantile(wierdX, 0.99)
	if !math.IsInf(infRes, 1) {
		t.Errorf("Expected inf, got: %f", infRes)
	}

	x := 43.12345
	constX := []float64{x, x, x, x, x, x, x, x, x, x, x, x}
	q := make([]float64, 5)
	q[0] = Quantile(constX, 0.0)
	q[1] = Quantile(constX, 0.14)
	q[2] = Quantile(constX, 0.54)
	q[3] = Quantile(constX, 0.95)
	q[4] = Quantile(constX, 1.0)

	for i, e := range q {
		if !closeEnough(x, e, 0.000001) {
			t.Errorf("Expected 43.12345, got: %f for i = %d\n", e, i)
		}
	}

	if !math.IsNaN(Quantile(constX, 2.5)) {
		t.Errorf("Expected NaN, got: %f \n", Quantile(constX, 2.5))
	}

	if !math.IsNaN(Quantile(constX, -10.0)) {
		t.Errorf("Expected NaN, got: %f \n", Quantile(constX, 2.5))
	}
}

// This unit test tests if Quantiles function is consistent with Quantile function.
func TestUnitQuantiles(t *testing.T) {
	epsilon := 0.0001
	input := []float64{14.54, 20.12, 50.0, 211.0, -10.5, 3.14159}
	probs := []float64{0.0, 1.0, -19.0, 44.123, 0.04, 0.000123, 0.995, 0.55, 0.75, 0.88}

	expected := make([]float64, len(probs))
	for i, p := range probs {
		expected[i] = Quantile(input, p)
	}
	result := Quantiles(input, probs)

	for i, res := range result {
		if !closeEnough(res, expected[i], epsilon) {
			t.Errorf("Quantile vs Quantiles failed for %d: Expected: %f, got: %f\n",
				i, expected[i], res)
		}
	}
}

func BenchmarkQuantile(b *testing.B) {
	input, _ := quantileResultsFromR2()
	var q float64

	for i := 0; i < b.N; i++ {
		q = Quantile(input, 0.90)
	}
	fmt.Sprintf("%f", q)
}

func BenchmarkQuantiles(b *testing.B) {
	input, _ := quantileResultsFromR2()
	p := make([]float64, 100)
	for i := 0; i < 100; i++ {
		p[i] = rand.Float64()
	}
	var q []float64

	for i := 0; i < b.N; i++ {
		q = Quantiles(input, p)
	}
	fmt.Sprintf("%v", q)
}

// Function quantileResultsFromR1 returns x slice and its quantiles calculated
// in R using quantile().
func quantileResultsFromR1() ([]float64, map[float64]float64) {
	input := []float64{14.54, 20.12, 50.0, 211.0, -10.5, 3.14159}
	res := make(map[float64]float64)
	res[0.0] = -10.50
	res[0.05] = -7.08960
	res[0.15] = -0.2688075
	res[0.55] = 18.7250
	res[0.75] = 42.530
	res[0.90] = 130.500
	res[0.99] = 202.950
	res[1.00] = 211.0

	return input, res
}

func quantileResultsFromR2() ([]float64, map[float64]float64) {
	input := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = float64(i + 1)
	}

	res := make(map[float64]float64)
	res[0.0] = 1.0
	res[0.05] = 50.95
	res[0.15] = 150.85
	res[0.55] = 550.45
	res[0.75] = 750.25
	res[0.90] = 900.10
	res[0.99] = 990.01
	res[1.00] = 1000.0

	return input, res
}

func closeEnough(x, y, epsilon float64) bool {
	if math.IsNaN(x) && math.IsNaN(y) {
		return true
	}
	if math.IsInf(x, 1) && math.IsInf(y, 1) {
		return true
	}
	if math.IsInf(x, -1) && math.IsInf(y, -1) {
		return true
	}
	if math.Abs(x-y) < epsilon {
		return true
	}
	return false
}
