package discrete

import "testing"

func TestFactorial(t *testing.T) {

	nlist := []int64{8, -1, 3, 0, 7, 4, -5, -4, -20, 2}
	for _, n := range nlist {
		f, err := Factorial(n)
		if err != nil {
			if isNegativeInteger(n) {
				t.Logf("Factorial(%d) = %d success catching error: %s", n, f, err)
			} else {
				t.Errorf("Factorial(%d) = %d error: %s", n, f, err)
			}
		} else {
			t.Logf("Factorial(%d) = %d success", n, f)
		}
	}

	type inout struct {
		in  int64
		out int64
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

func TestTotalPermutations(t *testing.T) {

	type permtest struct {
		n int64
		r int64
		p int64
	}

	permlist := []permtest{{10, 7, 604800}, {6, 6, 720}, {5, 4, 120}}
	for _, n := range permlist {
		p, err := TotalPermutations(n.n, n.r)
		if err != nil {
			t.Errorf("Permputation(%d, %d) = %d error: %s", n.n, n.r, p, err)
		} else if p == n.p {
			t.Logf("Permputation(%d, %d) = %d success", n.n, n.r, n.p)
		} else {
			t.Errorf("Permputation(%d, %d) = %d should be %d", n.n, n.r, n.p, p)
		}

	}

}

func TestPolynomial(t *testing.T) {

	type testValues struct {
		x float32
		n uint8
		p float64
	}

	successValues := []testValues{{1, 2, 1}, {2, 3, 8}}

	for _, v := range successValues {
		p := Polynomial(v.x, v.n)
		if p == v.p {
			t.Logf("pass: func Polynomial(%v, %v) -> %v == %v", v.x, v.n, p, v.p)
		} else {
			t.Errorf("fail: func Polynomial(%v, %v) -> %v != %v", v.x, v.n, p, v.p)
		}
	}

}
