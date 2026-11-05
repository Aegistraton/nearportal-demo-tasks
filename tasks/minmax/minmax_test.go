package minmax

import (
	"reflect"
	"testing"
)

func TestMinMax(t *testing.T) {
	cases := []struct {
		values   []int
		min, max int
	}{
		{[]int{7}, 7, 7}, {[]int{3, 1, 9, 2}, 1, 9}, {[]int{-8, -2, -15}, -15, -2}, {[]int{4, 4, 4}, 4, 4}, {[]int{-3, 0, 12, -3}, -3, 12},
	}
	for _, tc := range cases {
		input := append([]int(nil), tc.values...)
		before := append([]int(nil), input...)
		min, max, err := MinMax(input)
		if err != nil || min != tc.min || max != tc.max {
			t.Errorf("MinMax(%v) = %d, %d, %v; want %d, %d, nil", tc.values, min, max, err, tc.min, tc.max)
		}
		if !reflect.DeepEqual(input, before) {
			t.Errorf("input modified: before %v, after %v", before, input)
		}
	}
}
func TestMinMaxEmpty(t *testing.T) {
	for _, input := range [][]int{nil, {}} {
		min, max, err := MinMax(input)
		if err == nil || min != 0 || max != 0 {
			t.Errorf("empty input: got %d, %d, %v; want 0, 0 and an error", min, max, err)
		}
	}
}
