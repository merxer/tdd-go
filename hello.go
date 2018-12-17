package main

import "fmt"

const helloPrefix = "Hello, "
const spanish = "Spanish"
const thai = "Thai"
const spanishHelloPrefix = "Hola, "
const thaiHelloPrefix = "สวัสดีครับ, "

// Hello  return string
func Hello(name, language string) (result string) {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == thai {
		return thaiHelloPrefix + name
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Pat", "spanish"))
}
