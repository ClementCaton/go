package main

func main() {
}

func greeter(name string, lang string) string {
	switch lang {
	case "french":
		return "Bonjour " + name
	case "spanish":
		return "Hola " + name
	default:
		return "Hello " + name
	}
}
