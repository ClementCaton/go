package main

import "testing"

func TestPrinter(t *testing.T) {
	var a = printer("Hello World!")
	if a == "Hello World!" {
		return
	} else {
		t.Error("Expected 'Hello World!', Got ", a)
	}
}
