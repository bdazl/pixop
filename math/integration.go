package math

// Definite integral of f over [a, b].

func DefiniteIntegration(a, b float64, stepCount int, f func(float64) float64) float64 {
	// Trapezoidal integration
	n := stepCount

	ba := b - a
	ban := ba / float64(n)

	sum := f(a)/2 + f(b)/2
	for k := 1; k < n; k++ {
		sum += f(a + float64(k)*ban)
	}

	return ban * sum
}
