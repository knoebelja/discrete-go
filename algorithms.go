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

	if !(n >= r) {
		err = fmt.Errorf("error: parameters n %d >= r %d != true", n, r)
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

// Power computes p = x^n given float64 x and uint n,
func Power(x float32, n uint) (s float64) {
	xf64 := float64(x)
	s = xf64

	for k := uint(1); k < n; k++ {
		s *= xf64
	}
	return
}

// NextSubset given a string of 0's and 1's representing a subset, gives the next subset from 000...0 to 111...1
func NextSubset(a string) (z string, err error) {
	for _, r := range a {
		if r != '0' && r != '1' {
			err = fmt.Errorf("error: parameter a string contains characters that are not 0 or 1")
			return
		}
	}

	n := len(a)
	if n == 0 {
		err = fmt.Errorf("error: parameter a string must have length greater than 0")
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

// Compare determines the validity of an asserted relationship between 2 int values.
// ie. If F(x, y) returns x < y then F(1, 2) returns true; F(2, 1) returns false.
type Compare func(x, y int) (ok bool)

// BubbleSort mutates s []int so that for all s[0], s[1], ..., s[n-1], s[n] Compare(s[k-1], s[k]) is true
func BubbleSort(s []int, c Compare) {
	n := len(s) - 1
	for j := 0; j < n; j++ {
		for k := n - 1; k >= j; k-- {
			if !c(s[k], s[k+1]) {
				placeholder := s[k+1]
				s[k+1] = s[k]
				s[k] = placeholder
			}
		}
	}
	return
}
