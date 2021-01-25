package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello takes one string and returns the greeting string
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}