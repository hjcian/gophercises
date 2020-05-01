package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const engPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = engPrefix
	}
	return
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func main() {
	fmt.Println(Hello("", ""))
}
