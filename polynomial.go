package discrete

import "fmt"

// Polynomial represents a list of constants (a[0], a[1], ..., a[n])
// in the algorith P(x) = a[0]*x^0 + a[1]*x^1 + ... +a[n]*x^n
// where n equals the length of Polynomial minus 1 and x is any number
type Polynomial []float64

// NewPolynomial returns
func NewPolynomial(constants ...float64) (p Polynomial, err error) {
	if len(constants) == 0 {
		err = fmt.Errorf("error: at least one parameter must be set")
		p = nil
		return
	}

	p = constants
	return
}

func (p Polynomial) String() (s string) {
	for i := range p {
		curr := p[i]
		if curr < 0 {
			if len(s) > 0 {
				s += " - "
			} else {
				s += "-"
			}
			curr *= -1
		} else if curr > 0 {
			if len(s) > 0 {
				s += " + "
			}
		}
		s += fmt.Sprintf("%v", curr)
		if i > 0 {
			s += "x"
		}
		if i > 1 {
			s += fmt.Sprintf("^%d", i)
		}
	}
	return
}

// BruteEval computes a brute force evaluation of p Polynomial by finding the sum of all p[i]*x^i where i is an index of p
func (p Polynomial) BruteEval(x float64) (sum float64) {
	for k := 0; k < len(p); k++ {
		sum += float64(p[k]) * Power(x, uint(k))
	}
	return
}

// Eval computes a faster evaluation of p Polynomial using Horner's algorithm
func (p Polynomial) Eval(x float64) (sum float64) {
	xf64 := float64(x)
	n := len(p) - 1
	sum = float64(p[n])
	for k := 1; k <= n; k++ {
		sum = sum*xf64 + float64(p[n-k])
	}
	return
}
