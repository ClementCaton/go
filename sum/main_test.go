package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		b    int
		c    int
		want int
	}{
		{1, 1, 2},
		{3, 1, 4},
		{2, 2, 4},
		{5, 5, 10},
		{10, 3, 13},
		{0, 0, 0},
		{9, 56, 65},
		{100, 100, 200},
		{1024, 2048, 3072},
	}

	for _, test := range cases {
		t.Run("sum", func(t *testing.T) {
			var a = sum(test.b, test.c)
			if a != test.want {
				t.Errorf("Expected 5, got %d", a)
			}
		})
	}
}
