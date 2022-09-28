package main

import "testing"

//test printer function
func TestPrinter(t *testing.T) {
	var a = greeter("Algosup")
	var want = "Hello Algosup"
	if a != want {
		t.Errorf("Expected %q, Got %q", want, a)
	}
}
