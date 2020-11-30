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

// Permutations calcutes p permutations of n objects, taken r at a time: n! / (n - r)!
func Permutations(n int64, r int64) (p int64, err error) {
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

	if n-r == 0 || n-r == 1 {
		p, err = Factorial(n)
		return
	}

	p = func() (p int64) {

		dividends := make([]int64, 0)

		dividerMax := n - r

		for i := int64(n); i > dividerMax; i-- {
			dividends = append(dividends, i)
		}

		p = 1
		for _, dividend := range dividends {
			p *= dividend
		}
		return
	}()
	return
}
