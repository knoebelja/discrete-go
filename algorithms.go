package discrete

import (
	"fmt"
)

// Factorial calcultes n! which equals 1 * 2 * 3 * ... * n
func Factorial(n int64) (f int64, err error) {
	if err = checkNonnegativeInteger("checking parameter n", n); err != nil {
		return
	}

	f = 1
	for i := int64(2); i <= n; i++ {

		f *= i
	}
	return
}

// TotalPermutations calcutes p permutations of n objects, taken r at a time: n! / (n - r)!
func TotalPermutations(n, r int64) (p int64, err error) {

	if err = checkNonnegativeInteger("checking parameter n", n); err != nil {
		return
	}

	if err = checkNonnegativeInteger("checking parameter r", r); err != nil {
		return
	}

	if r > n {
		err = fmt.Errorf("comparing parameters: n{%d} must be greater than r{%d}", n, r)
		return
	}

	if r == 0 {
		p = 1
		return
	}

	dividerMax := n - r
	if dividerMax == 0 || dividerMax == 1 {
		p, err = Factorial(n)
		return
	}

	p = 1
	for i := n; i > dividerMax; i-- {
		p *= i
	}

	return
}

// Polynomial computes p = x^n given float64 x and uint32 n,
func Polynomial(x float32, n uint8) (p float64) {
	xf64 := float64(x)
	p = xf64

	for k := uint8(1); k < n; k++ {
		p *= xf64
	}
	return
}

// EvaluatePolynomial computesa float64 A*x^n, where A is a Set of uint32, n is a uint32 and x is the length of Set A
func EvaluatePolynomial(a []uint32, x float32) (s float64, err error) {
	if len(a) == 0 {
		err = fmt.Errorf("parameter a []uint32 must be of length greater than 0")
		return
	}

	s = float64(a[0])
	for k := 1; k < len(a); k++ {
		s += float64(float32(a[k]) * x)
	}
	return
}
