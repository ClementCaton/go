package main

import "os"

func main() {
	var input string = os.Args[1]
	var a = greeter(input, "french")
	println(a)
}

func greeter(input string, lang string) string {
	if lang == "french" {
		return "Bonjour " + input
	}
	return "Hello " + input
}
