package main

import "fmt"

// Hello ...
func Hello(name, lang string) string {
	switch lang {
	case "Spanish":
		return "Hola, " + name
	default:
		if name == "" {
			return "Hello, World"
		}
		return "Hello, " + name
	}
}

// GoodBye ...
func GoodBye(name string) string {
	name = fmt.Sprintf("Goodbye %s, sorry to see you leave!", name)
	return name
}

func main() {
	fmt.Println(Hello("Luke"))
}
