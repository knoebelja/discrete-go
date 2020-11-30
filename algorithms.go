package discrete

import (
	"errors"
	"fmt"
	"reflect"
)

func Factorial(n NonnegativeInteger) (f PositiveInteger, err error) {
	if err = Validate(n); err != nil {
		return
	}

	f = 1

	pn := PositiveInteger(n)
	if pn.IsValid() {

		for i := PositiveInteger(2); i <= pn; i++ {

			f *= i
		}
	}

	err = Validate(f)
	return
}

func Permutations(n NonnegativeInteger, r NonnegativeInteger) (p PositiveInteger, err error) {
	if err = Validate(n, r); err != nil {
		return
	}

	if r > n {
		err = errors.New(fmt.Sprintf("%d > %d", r, n))
		return
	}

	if r == 0 {
		p = 1
		return
	}

	p1, err := Factorial(n)
	if err != nil {
		return
	}

	if r == n || r == n-1 {
		p = p1
		return
	}

	p2, err := Factorial(n - r)
	if err != nil {
		return
	}

	p = p1 / p2
	err = Validate(p)
	return
}

func IsElementOf(e Element, s Set) (ok bool) {

	for _, se := range s {
		if ok = reflect.ValueOf(se) == reflect.ValueOf(e); ok {
			return
		}
	}

	return
}
