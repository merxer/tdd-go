package main

import "fmt"

const helloPrefix = "Hello, "

// Hello  return string
func Hello(s string) (result string) {
	if s != "" {
		result = helloPrefix + s
		return
	}
	result = helloPrefix + "world"
	return
}

func main() {
	fmt.Println(Hello("world"))
}
