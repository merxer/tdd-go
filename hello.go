package main

import "fmt"

const helloPrefix = "Hello, "

// Hello  return string
func Hello(name, language string) (result string) {
	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		helloPrefix := "Hola, "
		return helloPrefix + name
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Pat", "spanish"))
}
