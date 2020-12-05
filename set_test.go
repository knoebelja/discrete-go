package discrete

import (
	"testing"
)

func TestIsElementOf(t *testing.T) {
	type params struct {
		e interface{}
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
			t.Logf("success: %#v is element of %#v", s.e, s.s)
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

func TestElements(t *testing.T) {

	set := New("4", 5, 2123123, 5234, 4234, 34, 34, 42, 34, 423, 42, 34, 234, "a")

	feelslikeinfinity := 999999999999999999
	equal := true
	for equal && feelslikeinfinity > 0 {

		e1 := ListElements(set)
		e2 := ListElements(set)

		if len(e1) != len(e2) {
			t.Fatalf("fail: length of %v is not equal to length of %v", e1, e2)
		}

		for i := 0; i < len(e1); i++ {
			if e1[i] != e2[i] {
				equal = false
				break
			}

		}

		if equal {
			t.Errorf("fail: order of %v == %v", e1, e2)
		} else {
			t.Logf("pass: order of %v != %v", e1, e2)
		}

	}

	if equal {
		t.Fatalf("fail: ListElements(%v) always equals %v", set, ListElements(set))
	}

}
