package stat

import (
	"math"
	"testing"
)

func TestUnitMean(t *testing.T) {
	res := make([]float64, 3)
	expected := make([]float64, 3)

	res[0] = Mean([]float64{1.0, 1.0, 1.0, 1.0, 1.0})
	res[1] = Mean([]float64{99.5})
	res[2] = Mean([]float64{-10.0, 10.0})

	expected[0] = 1.0
	expected[1] = 99.5
	expected[2] = 0.0

	for i := 0; i < 3; i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected: %f | Got: %f", expected[i], res[i])
		}
	}
}

func TestUnitSum(t *testing.T) {
	res := make([]float64, 4)
	expected := make([]float64, 4)

	res[0] = Sum([]float64{1.0, 1.0, 1.0, 1.0, 1.0})
	res[1] = Sum([]float64{99.5})
	res[2] = Sum([]float64{})
	res[3] = Sum([]float64{11111111.1, 11111111.1})

	expected[0] = 5.0
	expected[1] = 99.5
	expected[2] = 0.0
	expected[3] = 22222222.2

	for i := 0; i < 3; i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected: %f | Got: %f", expected[i], res[i])
		}
	}
}

func TestUnitVariance(t *testing.T) {
	input := []float64{15.123, 5.321, 10.123, -166.100, 321.123}
	resultInR := 31044.99
	goVar := Variance(input)

	if math.Abs(goVar-resultInR) > 0.01 {
		t.Errorf("Expected variance = %f | Got: %f", resultInR, goVar)
	}
}
