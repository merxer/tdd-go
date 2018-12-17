package main

import "fmt"

const helloPrefix = "Hello, "

// Hello  return string
func Hello(name string) (result string) {
	if name == "" {
		name = "world"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
