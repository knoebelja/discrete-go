package discrete

import "testing"

func TestFactorial(t *testing.T) {

	type io struct{ in, out uint }

	values := []io{{0, 1}, {1, 1}, {2, 2}, {3, 6}, {4, 24}, {5, 120}, {6, 720}, {9, 362880}, {10, 3628800}}
	for _, n := range values {
		if f := Factorial(n.in); f != n.out {
			t.Errorf("fail: Factorial(%d) -> %d != %d", n.in, f, n.out)
		} else {
			t.Logf("pass: Factorial(%d) -> %d == %d", n.in, f, n.out)
		}
	}
}

func TestTotalPermutations(t *testing.T) {

	type io struct{ n, r, p uint }

	values := []io{{10, 7, 604800}, {6, 6, 720}, {5, 4, 120}}
	for _, n := range values {
		p, err := TotalPermutations(n.n, n.r)
		if err != nil {
			t.Errorf("fail: TotalPermutations(%d, %d) -> %s != nil", n.n, n.r, err)
		} else if p == n.p {
			t.Logf("pass: TotalPermputation(%d, %d) -> %d == %d", n.n, n.r, p, n.p)
		} else {
			t.Errorf("fail: TotalPermputation(%d, %d) -> %d != %d", n.n, n.r, p, n.p)
		}

	}

}

func TestExponential(t *testing.T) {

	type io struct {
		x float32
		n uint
		p float64
	}

	values := []io{{1, 2, 1}, {2, 3, 8}}

	for _, v := range values {
		p := Exponential(v.x, v.n)
		if p == v.p {
			t.Logf("pass: Polynomial(%v, %v) -> %v == %v", v.x, v.n, p, v.p)
		} else {
			t.Errorf("fail: Polynomial(%v, %v) -> %v != %v", v.x, v.n, p, v.p)
		}
	}

}

// func TestPolynomial(t *testing.T) {

// }

func TestNextSubset(t *testing.T) {

	type io struct{ in, out string }
	values := []io{{"0", "1"}, {"1001", "1010"}, {"010", "011"}, {"10011", "10100"}, {"100", "101"}, {"1110", "1111"}, {"111", ""}}

	for _, v := range values {
		z, err := NextSubset(v.in)
		if err != nil {
			t.Errorf("fail: NextSubset(%s) -> %d != nil", v.in, err)
		} else if v.out != z {
			t.Logf("fail: NextSubset(%s) -> %s != %s", v.in, v.out, z)
		} else {
			t.Logf("pass: NextSubset(%s) -> %s == %s", v.in, v.out, z)
		}
	}

}

func TestBubbleSort(t *testing.T) {

	c1 := Compare(func(x, y int) bool {
		return x > y
	})

	c2 := Compare(func(x, y int) bool {
		return x < y
	})

	s := []int{0, 5, 8, 2, -7, 8, 12}
	t.Logf("%v", s)
	BubbleSort(s, c1)
	t.Logf("%v", s)
	BubbleSort(s, c2)
	t.Logf("%v", s)

}
