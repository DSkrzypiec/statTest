package statTest

// Function mean calculates mean of given float64 vector.
func mean(x []float64) float64 {
	if len(x) == 0 {
		return 0.0
	}
	m := 0.0

	for _, e := range x {
		m += e
	}

	return m / float64(len(x))
}
