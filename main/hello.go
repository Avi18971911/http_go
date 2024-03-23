package main

func Hello(greeting string) string {
	if greeting == "" {
		greeting = "World"
	}
	return "Hello, " + greeting
}
