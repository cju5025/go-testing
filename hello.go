package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Colter"))
}

func Hello(name string) string {
	return "Hello, " + name + "."
}
