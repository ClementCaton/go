package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		b, c := 2, 3
		var a = sum(b, c)
		if a != b+c {
			t.Errorf("Expected %q, Got %q", b+c, a)
		}
	})
}
