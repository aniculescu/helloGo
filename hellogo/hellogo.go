package main

import "fmt"

const englishGreetingPrefix = "Greeting, "

func Hello(n string) string {
	if n == "" {
		n = "default"
	}
	return englishGreetingPrefix + n
}

func main() {
	fmt.Println(Hello("name"))
}
