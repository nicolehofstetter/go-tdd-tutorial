package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello"
const spanishPrefix = "Hola"
const frenchPrefix = "Bonjour"

func Hello(name string, language string) string {

	prefix := englishHelloPrefix

	if name == "" {
		name = "World"
	}

	if language == spanish {
		prefix = spanishPrefix
	}

	if language == french {
		prefix = frenchPrefix
	}

	return prefix + " " + name
}

func main() {
	fmt.Println(Hello("", ""))
}
