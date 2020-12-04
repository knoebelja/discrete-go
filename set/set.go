package set

import "reflect"

// Set is an unordered list of unique Elements
type Set *map[interface{}]bool

// New returns a Set of unique Elements
func New(elements ...interface{}) (s Set) {
	set := make(map[interface{}]bool)
	s = Set(&set)

	for _, e := range elements {
		Add(e, s)
	}

	return
}

// IsElementOf checks if Element E is in Set S
func IsElementOf(e interface{}, s Set) (ok bool) {

	for x := range *s {

		if reflect.DeepEqual(x, e) {
			ok = true
			return
		}

		es, eok := e.(Set)
		xs, xok := x.(Set)
		if eok && xok {
			if Equals(es, xs) {
				ok = true
				return
			}
		}

	}

	return

}

// Contains checks if Set A contains Set B, which means every Element of Set B is an Element of Set A
func Contains(a, b Set) (ok bool) {

	for x := range *a {
		if !IsElementOf(x, b) {
			return
		}
	}

	ok = true
	return
}

// Equals checks if Set A contains Set B and Set B contains Set A
func Equals(a, b Set) (ok bool) {
	ok = Contains(a, b) && Contains(b, a)
	return
}

// Cardinal returns the total number Elements in Set S
func Cardinal(s Set) (c int) {
	c = len(*s)
	return
}

// Add adds the Element E to Set S if it does not contain E already
func Add(e interface{}, s Set) (ok bool) {
	if IsElementOf(e, s) {
		return
	}
	ok = true
	(*s)[e] = true
	return
}
