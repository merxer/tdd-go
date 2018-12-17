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

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case thai:
		prefix = thaiHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Pat", "spanish"))
}
