package main

import "fmt"

const englishGreetingPrefix = "Greeting, "
const spanish = "Spanish"
const french = "French"
const frenchGreetingPrefix = "Bonjour, "
const spanishGreetingPrefix = "Hola, "

func GreetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchGreetingPrefix
	case spanish:
		prefix = spanishGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}

	return
}
func Hello(n, lang string) string {
	if n == "" {
		n = "default"
	}

	return GreetingPrefix(lang) + n
}

func main() {
	fmt.Println(Hello("name", ""))
}
