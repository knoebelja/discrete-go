package discrete

import (
	"fmt"
)

func isPositiveInteger(i int64) (ok bool) {
	return i > 0
}

func isNegativeInteger(i int64) (ok bool) {
	return i < 0
}

func isNonpositiveInteger(i int64) (ok bool) {
	return i <= 0
}

func isNonnegativeInteger(i int64) (ok bool) {
	return i >= 0
}

func isEmptySet(s []interface{}) (ok bool) {
	return len(s) == 0
}

func check(ok bool, msg string) (err error) {
	if !ok {
		err = fmt.Errorf(msg)
		return
	}
	return
}

func checkPositiveInteger(errwhen string, i int64) (err error) {
	return check(isPositiveInteger(i), fmt.Sprintf("%s: %d must be greater than 0", errwhen, i))
}

func checkNegativeInteger(errwhen string, i int64) (err error) {
	return check(isNegativeInteger(i), fmt.Sprintf("%s: %d must be less than 0", errwhen, i))
}

func checkNonnegativeInteger(errwhen string, i int64) (err error) {
	return check(isNonnegativeInteger(i), fmt.Sprintf("%s: %d must be less than or equal to 0", errwhen, i))
}

func checkNonpositiveInteger(errwhen string, i int64) (err error) {
	return check(isNonpositiveInteger(i), fmt.Sprintf("%s: %d must be greater than or equal to 0", errwhen, i))
}
