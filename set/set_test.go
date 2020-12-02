package set

import (
	"testing"
)

func TestIsElementOf(t *testing.T) {
	type params struct {
		e Element
		s Set
	}

	successes := []params{
		{"a", New("a")},
		{struct{}{}, New(struct{}{})},
		{New(New("a")), New(New(New("a")))},
		{New(New(New("a"))), New(New(New(New("a"))))},
		{New(1, New(New("a"), "b")), New(New(New("b", New("a")), 1))},
		{-5, New("5", -5)},
	}

	for _, s := range successes {

		if ok := IsElementOf(s.e, s.s); ok {
			t.Logf("success: %#v is element of %#v", &s.e, &s.s)
		} else {
			t.Errorf("failure: %#v is element of %#v", s.e, s.s)
		}

	}

	failures := []params{
		{1, New("a")},
		{[]struct{}{}, New(struct{}{})},
		{New(New("a")), New(New(New("a", "b")))},
		{New(New(New("a"), New("b"))), New(New(New(New("a"))))},
		{New(New(New(1))), New(New(New(New("a"))))},
		{-5, New("-5", 5)},
	}

	for _, f := range failures {

		if ok := !IsElementOf(f.e, f.s); ok {
			t.Logf("success: %#v is not element of %#v", f.e, f.s)
		} else {
			t.Errorf("failure: %#v is not element of %#v", f.e, f.s)
		}

	}

}

func TestContains(t *testing.T) {}

func TestEquals(t *testing.T) {
	type params struct {
		a Set
		b Set
	}

	successes := []params{
		{New(), New()},
		{New(3, 2, 1), New(1, 2, 3)},
		{New("a", "a"), New("a")},
		{New("a", New("a"), New("a", New("b"))), New(New(New("b"), "a"), New("a"), "a")},
	}

	for _, s := range successes {

		if ok := Equals(s.a, s.b); ok {
			t.Logf("success %#v equals %#v", s.a, s.b)
		} else {
			t.Errorf("failure: %#v does not equal %#v", s.a, s.b)
		}

	}

	failures := []params{
		{New(), New(New())},
		{New("a"), New("b")},
		{New("a", New("a"), New("a", New("b"))), New(New(New("b"), "a"), New("a"), "z")},
		{New("a", New("a"), New("a", New("b"))), New(New(New("z"), "a"), New("a"), "a")},
	}

	for _, f := range failures {
		if ok := Equals(f.a, f.b); !ok {
			t.Logf("success: %#v does not equal %#v", f.a, f.b)
		} else {
			t.Errorf("failure:  %#v equals %#v", f.a, f.b)
		}

	}

}
