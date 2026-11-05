package minmax

import "errors"

// MinMax returns the smallest and largest values without modifying values.
func MinMax(values []int) (int, int, error) {
	if len(values) == 0 {
		return 0, 0, errors.New("cannot find min/max of an empty slice")
	}

	min, max := values[0], values[0]

	for _, value := range values[1:] {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max, nil
}