package discrete

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestFactorial(c *check.C) {

	type io struct{ n, f uint }

	values := []io{{0, 1}, {1, 1}, {2, 2}, {3, 6}, {4, 24}, {5, 120}, {6, 720}, {9, 362880}, {10, 3628800}}
	for _, v := range values {
		f := Factorial(v.n)
		c.Logf("Factorial(n: %d) => (f: %d)", v.n, f)
		c.Check(f, check.Equals, v.f)
	}
}

func (s *S) TestTotalPermutations(c *check.C) {

	type io struct{ n, r, p uint }

	values := []io{{10, 7, 604800}, {6, 6, 720}, {5, 4, 120}, {4, 5, 120}}
	for _, v := range values {
		p, err := TotalPermutations(v.n, v.r)
		c.Logf("TotalPermutations(n: %d, r: %d) => (p: %d, err: %+v)", v.n, v.r, p, err)
		if err != nil {
			c.Check(err, check.ErrorMatches, "error:.+")
		} else {
			c.Check(p, check.Equals, v.p)
		}
	}

}

func (s *S) TestPower(c *check.C) {

	type io struct {
		x float32
		n uint
		s float64
	}

	values := []io{{1, 2, 1}, {2, 3, 8}}
	for _, v := range values {
		s := Power(v.x, v.n)
		c.Logf("Power(x: %+v, n: %+v) => (s: %+v)", v.x, v.n, s)
		c.Check(v.s, check.Equals, s)
	}

}

func (s *S) TestNextSubset(c *check.C) {

	type io struct{ a, z string }
	values := []io{{"0", "1"}, {"1001", "1010"}, {"010", "011"}, {"10011", "10100"}, {"100", "101"}, {"1110", "1111"}, {"111", ""}, {"", ""}, {"01010BANANANA", ""}}

	for _, v := range values {
		z, err := NextSubset(v.a)
		c.Logf("NextSubset(a: %s) => (z: %s, err: %+v)", v.a, z, err)
		if err != nil {
			c.Check(err, check.ErrorMatches, "error: .+")
		} else {
			c.Check(z, check.Equals, v.z)
		}
	}

}

func (s *S) TestPrintNextSubset(c *check.C) {

	var subset string = "0000000"
	var err error
	subsets := make([]string, 0)
	subsets = append(subsets, subset)

	for subset != "" {
		subset, err = NextSubset(subset)
		if err != nil {
			c.Fatal(err)
		}
		subsets = append(subsets, subset)
	}

	c.Logf("subsets = %v", subsets)
}

func (s *S) TestBubbleSort(c *check.C) {

	sort := []int{0, 5, 8, 2, -7, 8, 12}
	c.Logf("no sort => %+v", sort)

	c1 := Compare(func(x0, x1 int) bool {
		return x0 > x1
	})

	BubbleSort(sort, c1)
	c.Logf("greater than => %+v", sort)
	greater := []int{12, 8, 8, 5, 2, 0, -7}
	c.Check(sort, check.DeepEquals, greater)

	c2 := Compare(func(x0, x1 int) bool {
		return x0 < x1
	})

	BubbleSort(sort, c2)
	c.Logf("less than => %+v", sort)
	less := []int{-7, 0, 2, 5, 8, 8, 12}
	c.Check(sort, check.DeepEquals, less)

}
