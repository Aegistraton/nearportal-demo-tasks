package factorial

import "errors"

// Factorial returns n! for 0 <= n <= 20.
// For n > 20, return 0 and a non-nil error.
func Factorial(n uint64) (uint64, error) {
	if n > 20 {
		return 0, errors.New("factorial: input overflows uint64 (max is 20)")
	}

	result := uint64(1)
	for i := uint64(2); i <= n; i++ {
		result *= i
	}

	return result, nil
}
