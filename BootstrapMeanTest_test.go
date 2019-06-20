package statTest

import (
	"fmt"
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

	fmt.Println("Result for similar means:")
	fmt.Printf("P-Value = %f \n", t1.PValue)

	fmt.Println("Result for very different means:")
	fmt.Printf("P-Value = %f \n", t2.PValue)

	fmt.Printf("Result for %d sampled floats: \n", n)
	fmt.Printf("P-Value = %f \n", t3.PValue)
}
