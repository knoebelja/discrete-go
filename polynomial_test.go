package discrete

import (
	"gopkg.in/check.v1"
)

type PolynomialSuite struct{}

var _ = check.Suite(&PolynomialSuite{})

func (s *PolynomialSuite) TestNewPolynomial(c *check.C) {

	p, err := NewPolynomial()
	c.Check(err, check.ErrorMatches, "error:.+")
	c.Check(p, check.DeepEquals, Polynomial(nil))
}

func (s *PolynomialSuite) TestString(c *check.C) {
	p, err := NewPolynomial(-29, 41, 5, -17, 35)
	c.Check(err, check.DeepEquals, error(nil))
	if c.Check(p.String(), check.Equals, "-29 + 41x + 5x^2 - 17x^3 + 35x^4") {
		c.Logf("P(x) = %s", p.String())
	}

	p, err = NewPolynomial(29, -41, -5, 17, -35)
	c.Check(err, check.DeepEquals, error(nil))
	if c.Check(p.String(), check.Equals, "29 - 41x - 5x^2 + 17x^3 - 35x^4") {
		c.Logf("P(x) = %s", p.String())
	}
}

func (s *PolynomialSuite) TestEval(c *check.C) {
	p, err := NewPolynomial(-29, 41, 5, -17, 35)
	c.Check(err, check.DeepEquals, error(nil))
	c.Check(len(p), check.Equals, 5)

	for x := float64(1); x <= 100; x++ {
		a := p.BruteEval(x)
		c.Logf("BruteEval(%v) => %v", x, a)
		b := p.Eval(x)
		c.Logf("     Eval(%v) => %v", x, b)
	}
}
