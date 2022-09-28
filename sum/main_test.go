package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		b := []int{1, 2, 3}
		var a = sum(b)
		if a != 6 {
			t.Errorf("Expected %d, Got %d", 6, a)
		}
	})
}
