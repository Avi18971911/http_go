package main

import "fmt"

func Hello(greeting string) string {
	if greeting == "" {
		greeting = "World"
	}
	return "Hello, " + greeting
}

func main() {
	fmt.Println(Hello("Mark"))
}
