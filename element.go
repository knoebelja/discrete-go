package discrete

import (
	"errors"
	"fmt"
)

type Element interface {
	IsValid() (ok bool)
	Type() (t string)
}

func Validate(objs ...Element) (err error) {
	for _, o := range objs {
		if !o.IsValid() {
			err = errors.New(fmt.Sprintf("Invalid %s: %#v", o.Type(), o))
			return
		}
	}
	return
}

type Integer int64

func (i Integer) IsValid() (ok bool) {
	ok = true
	return
}

func (i Integer) Type() (t string) {
	t = "Integer"
	return
}

type PositiveInteger int64

func (pi PositiveInteger) IsValid() (ok bool) {
	ok = pi > 0
	return
}

func (pi PositiveInteger) Type() (t string) {
	t = "Positive Integer"
	return
}

type NegativeInteger int64

func (ni NegativeInteger) IsValid() (ok bool) {
	ok = ni < 0
	return
}

func (ni NegativeInteger) Type() (t string) {
	t = "Negative Integer"
	return
}

type NonpositiveInteger int64

func (npi NonpositiveInteger) IsValid() (ok bool) {
	ok = npi <= 0
	return
}

func (npi NonpositiveInteger) Type() (t string) {
	t = "Nonpositive Integer"
	return
}

type NonnegativeInteger int64

func (nni NonnegativeInteger) IsValid() (ok bool) {
	ok = nni >= 0
	return
}

func (nni NonnegativeInteger) Type() (t string) {
	t = "Nonnegative Integer"
	return
}
