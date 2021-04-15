package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Colter"))
}

const greeting = "Hello, "

func Hello(name string) string {
	return greeting + name
}
