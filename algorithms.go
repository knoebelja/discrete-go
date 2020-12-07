package discrete

import "fmt"

// Factorial calcultes n! which equals 1 * 2 * 3 * ... * n
func Factorial(n uint) (f uint) {
	f = 1
	for i := uint(2); i <= n; i++ {

		f *= i
	}
	return
}

// TotalPermutations calcutes p permutations of n objects, taken r at a time: n! / (n - r)!
func TotalPermutations(n, r uint) (p uint, err error) {

	if r > n {
		err = fmt.Errorf("error: evaluating parameters r uint > n uint where %d > %d must be false", r, n)
		return
	}

	if r == 0 {
		p = 1
		return
	}

	dividerMax := n - r
	if dividerMax == 0 || dividerMax == 1 {
		p = Factorial(n)
		return
	}

	p = 1
	for i := n; i > dividerMax; i-- {
		p *= i
	}

	return
}

// Exponential computes p = x^n given float64 x and uint n,
func Exponential(x float32, n uint) (s float64) {
	xf64 := float64(x)
	s = xf64

	for k := uint(1); k < n; k++ {
		s *= xf64
	}
	return
}

// PolynomialEvaluation computes float64 A*x^n, where A is a Set of uint, n is a uint and x is the length of Set A
func PolynomialEvaluation(a []uint, x float32) (s float64, err error) {
	if len(a) == 0 {
		err = fmt.Errorf("error: a []uint must be of length greater than 0")
		return
	}

	xf64 := float64(x)
	s = float64(a[0])

	for k := 1; k < len(a); k++ {
		s += float64(a[k]) * xf64
	}
	return
}

// NextSubset given a string of 0's and 1's representing a subset, gives the next subset from 000...0 to 111...1
func NextSubset(a string) (z string, err error) {
	for _, r := range a {
		if r != '0' && r != '1' {
			err = fmt.Errorf("error: NextSubset(a string) accepts only a string of 0's and 1s")
			return
		}
	}

	n := len(a)
	if n == 0 {
		err = fmt.Errorf("error: NextSubset(a string) requires a string of length 1 or greater")
		return
	}

	k := n - 1
	for k >= 1 && a[k] == '1' {
		k--

	}

	if a[k] == '0' {

		for i := 0; i < n; i++ {
			if i == k {
				z += string('1')
			} else if i > k && a[i] == '1' {
				z += string('0')
			} else {
				z += string(a[i])
			}
		}
	}

	return
}
