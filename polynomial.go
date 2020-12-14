package discrete

import "fmt"

// Polynomial represents a list of constants (a[0], a[1], ..., a[n])
// in the algorith P(x) = a[0]*x^0 + a[1]*x^1 + ... +a[n]*x^n
// where n equals the length of Polynomial minus 1 and x is any number
type Polynomial []float32

// NewPolynomial returns
func NewPolynomial(constants ...float32) (p Polynomial, err error) {
	if len(constants) == 0 {
		err = fmt.Errorf("at least one parameter must be set")
		p = nil
		return
	}

	p = constants
	return
}

// BruteEval computes a brute force evaluation of p Polynomial by finding the sum of all p[i]*x^i where i is an index of p
func (p Polynomial) BruteEval(x float32) (sum float64) {
	for k := 0; k < len(p); k++ {
		sum += float64(p[k]) * Power(x, uint(k))
	}
	return
}

// Eval computes a faster evaluation of p Polynomial using Horner's algorithm
func (p Polynomial) Eval(x float32) (sum float64) {
	xf64 := float64(x)
	sum = float64(p[0])
	for k := 1; k < len(p); k++ {
		sum += float64(p[k]) * xf64
	}
	return
}
