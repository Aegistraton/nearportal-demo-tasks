package factorial

import "errors"

// Factorial returns n! for 0 <= n <= 20.
func Factorial(n uint64) (uint64, error) {
	if n > 20 {
		return 0, errors.New("factorial overflows uint64")
	}

	result := uint64(1)
	for i := uint64(2); i <= n; i++ {
		result *= i
	}

	return result, nil
}