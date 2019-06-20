package statTest

import (
	"testing"
)

func TestUnitMean(t *testing.T) {
	res := make([]float64, 3)
	expected := make([]float64, 3)

	res[0] = mean([]float64{1.0, 1.0, 1.0, 1.0, 1.0})
	res[1] = mean([]float64{99.5})
	res[2] = mean([]float64{-10.0, 10.0})

	expected[0] = 1.0
	expected[1] = 99.5
	expected[2] = 0.0

	for i := 0; i < 3; i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected: %f | Got: %f", expected[i], res[i])
		}
	}
}
