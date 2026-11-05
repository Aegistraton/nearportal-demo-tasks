package factorial

import "testing"

func TestFactorial(t *testing.T) {
	cases := []struct{ n, want uint64 }{{0, 1}, {1, 1}, {2, 2}, {3, 6}, {5, 120}, {10, 3628800}, {20, 2432902008176640000}}
	for _, tc := range cases {
		got, err := Factorial(tc.n)
		if err != nil || got != tc.want {
			t.Errorf("Factorial(%d) = %d, %v; want %d, nil", tc.n, got, err, tc.want)
		}
	}
}
func TestFactorialOverflow(t *testing.T) {
	for _, n := range []uint64{21, 100, ^uint64(0)} {
		got, err := Factorial(n)
		if err == nil || got != 0 {
			t.Errorf("Factorial(%d) = %d, %v; want 0 and an error", n, got, err)
		}
	}
}
