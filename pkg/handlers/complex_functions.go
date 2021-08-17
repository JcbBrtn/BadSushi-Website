package handlers

import "math"

func Mandlebrot(x complex128, c complex128) complex128 {
	return (x * x) + c
}

func Magnitude(x complex128) float64 {
	return math.Sqrt(real(x)*real(x) + imag(x)*imag(x))
}
