package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Colter"))
	fmt.Println(Hello(""))
}

const greeting = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greeting + name
}
