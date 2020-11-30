package discrete

// type Set []interface {
// 	Size() (s NonnegativeInteger)
// 	IsValid() (ok bool)
// 	Type() (t string)
// }

// type ElementSet []Element

// func (es ElementSet) IsValid() (ok bool) {

// 	check := make(ElementSet, 0)

// 	for _, e := range es {
// 		if !e.IsValid() {
// 			return
// 		}

// 		if ok = IsElementOf(e, check); !ok {
// 			return
// 		}

// 		check = append(check, e)
// 	}

// 	ok = true
// 	return
// }

// func (es ElementSet) Type() (t string) {
// 	t = "Element Set"
// 	return
// }

// func (es ElementSet) Size() (s NonnegativeInteger) {
// 	s = NonnegativeInteger(len(es))
// 	return
// }
