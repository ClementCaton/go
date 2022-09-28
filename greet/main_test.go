package main

import "testing"

//test printer function
func TestPrinter(t *testing.T) {
	t.Run("greet algosup", func(t *testing.T) {
		var a = greeter("Algosup")
		var want = "Hello Algosup"
		if a != want {
			t.Errorf("Expected %q, Got %q", want, a)
		}
	})
	t.Run("greet Lucien", func(t *testing.T) {
		var a = greeter("Lucien")
		var want = "Hello Lucien"
		if a != want {
			t.Errorf("Expected %q, Got %q", want, a)
		}
	})
}
