package main

import "fmt"

func Hello(n string) string {
	return "hello " + n
}

func main() {
	fmt.Println(Hello("name"))
}
