package statTest

import (
	"math/rand"
	"testing"
)

func TestUnitBootstrapMean(t *testing.T) {
	x1 := []float64{0.25, 0.50, 0.50, 0.25, 0.50, 0.75, 0.05}
	y1 := []float64{0.25, 0.50, 0.25, 0.50, 0.75}
	x2 := []float64{12313.123, 123.321, 0.123, 123.123, 4341.123}

	const n = 100000
	x3 := make([]float64, n)
	y3 := make([]float64, n)

	for i := 0; i < n; i++ {
		x3[i] = rand.Float64() * 100.0
		y3[i] = rand.Float64() * 100.0
	}

	t1 := BootstrapMean(x1, y1, 1000)
	t2 := BootstrapMean(x2, y1, 1000)
	t3 := BootstrapMean(x3, y3, 100)

	if t1.PValue <= 0.10 {
		t.Errorf("Expected P-Value > 0.10, got: %f\n", t1.PValue)
	}

	if t2.PValue > 0.10 {
		t.Errorf("Expected P-Value <= 0.10, got: %f \n", t2.PValue)
	}

	if t3.PValue <= 0.10 {
		t.Errorf("Expected P-Value > 0.10, got: %f\n", t3.PValue)
	}
}

func TestUnitBootstrapMeanSingle(t *testing.T) {
	const n = 10000
	x1 := make([]float64, n)

	for i := 0; i < n; i++ {
		x1[i] = rand.Float64() * 55.0
	}

	t1 := BootstrapMeanSingle(x1, 53.50, 0.05, 1000)
	t2 := BootstrapMeanSingle(x1, 10.0, 0.05, 1000)

	if t1.PValue <= 0.10 {
		t.Errorf("Expected P-Value > 0.10, got: %f\n", t1.PValue)
	}

	if t2.PValue > 0.10 {
		t.Errorf("Expected P-Value <= 0.10, got: %f \n", t2.PValue)
	}
}
