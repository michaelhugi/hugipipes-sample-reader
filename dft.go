package hugipipes_sample

import (
	"math"
)

func dftNative(input []float64) []complex128 {
	c := make([]complex128, len(input))
	arg := -2.0 * math.Pi / float64(len(input))
	for k := 0; k < len(input); k++ {
		r, i := 0.0, 0.0
		for n := 0; n < len(input); n++ {
			r += input[n] * math.Cos(arg*float64(n)*float64(k))
			i += input[n] * math.Sin(arg*float64(n)*float64(k))
		}
		c[k] = complex(r, i)
	}
	return c
}

func amplitude(c []complex128) []float64 {
	amp := make([]float64, len(c))
	for i := 0; i < len(c); i++ {
		amp[i] = math.Sqrt(real(c[i])*real(c[i]) + imag(c[i])*imag(c[i]))
	}
	return amp
}
