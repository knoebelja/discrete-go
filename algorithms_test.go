package discrete

import "testing"

func TestFactorial(t *testing.T) {

	nlist := []NonnegativeInteger{8, -1, 3, 0, 7, 4, -5, -4, -20, 2}
	for _, n := range nlist {
		f, err := Factorial(n)
		if err != nil {
			if NegativeInteger(n).IsValid() {
				t.Logf("Factorial(%d) = %d success catching error: %s", n, f, err)
			} else {
				t.Errorf("Factorial(%d) = %d error: %s", n, f, err)
			}
		} else {
			t.Logf("Factorial(%d) = %d success", n, f)
		}
	}

	type inout struct {
		in  NonnegativeInteger
		out PositiveInteger
	}

	iolist := []inout{{5, 120}, {6, 720}, {9, 362880}, {10, 3628800}}
	for _, n := range iolist {
		f, err := Factorial(n.in)
		if err != nil {
			t.Errorf("Factorial(%d) = %d error: %s", n, f, err)
		} else if f != n.out {
			t.Errorf("Factorial(%d) = %d should be %d", n.in, f, n.out)
		} else {
			t.Logf("Factorial(%d) = %d success", n.in, f)
		}
	}
}

func TestPermutations(t *testing.T) {

	type permtest struct {
		n NonnegativeInteger
		r NonnegativeInteger
		p PositiveInteger
	}

	permlist := []permtest{{10, 7, 604800}, {6, 6, 720}, {5, 4, 120}}
	for _, n := range permlist {
		p, err := Permutations(n.n, n.r)
		if err != nil {
			t.Errorf("P(%d, %d) = %d error: %s", n.n, n.r, p, err)
		} else if p == n.p {
			t.Logf("P(%d, %d) = %d success", n.n, n.r, n.p)
		} else {
			t.Errorf("P(%d, %d) = %d should be %d", n.n, n.r, n.p, p)
		}

	}

}
