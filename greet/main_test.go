package main

import "testing"

//test printer function
func TestPrinter(t *testing.T) {
	var a = greeter("Hello World!")
	if a != "Hello World!" {
		t.Error("Expected 'Hello World!', Got ", a)
	}
}
