package main

import "fmt"

// Hello  return string
func Hello(s string) string {
	result := "Hello, " + s
	return result
}

func main() {
	fmt.Println(Hello("world"))
}
