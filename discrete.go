package discrete

type Variable interface {
	IsValid() (valid bool)
}

func Validate(vars ...Variable) (valid bool) {
	for _, v := range vars {
		if !v.IsValid() {
			return
		}
	}

	valid = true
	return
}

type Integer int

func (i Integer) IsValid() (valid bool) {
	valid = true
	return
}

type PositiveInteger int

func (pi PositiveInteger) IsValid() (valid bool) {
	valid = pi > 0
	return
}

type NegativeInteger int

func (ni NegativeInteger) IsValid() (valid bool) {
	valid = ni < 0
	return
}

type NonpositiveInteger int

func (pi NonpositiveInteger) IsValid() (valid bool) {
	valid = pi <= 0
	return
}

type NonnegativeInteger int

func (ni NonnegativeInteger) IsValid() (valid bool) {
	valid = ni >= 0
	return
}
