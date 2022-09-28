package main

import "os"

func main() {
	var input string = os.Args[1]
	var a = greeter(input)
	println(a)
}

func greeter(input string) string {
	return "hello " + input
}
