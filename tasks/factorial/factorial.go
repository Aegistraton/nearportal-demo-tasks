package factorial

import "errors"

// Factorial returns n! for 0 <= n <= 20.
// For n > 20, return 0 and a non-nil error.
func Factorial(n uint64) (uint64, error) {
	return 0, errors.New("not implemented")
}
