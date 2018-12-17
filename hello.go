package main

import "fmt"

const helloPrefix = "Hello, "

// Hello  return string
func Hello(s string) string {
	result := helloPrefix + s
	return result
}

func main() {
	fmt.Println(Hello("world"))
}
