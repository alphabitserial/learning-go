package main

import "fmt"

func main() {
	fmt.Println(Hello("Nadia", ""))
}

const spanish = "Spanish"
const french = "French"
const mandarin = "Mandarin"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const mandarinHelloPrefix = "Ni hao, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case mandarin:
		prefix = mandarinHelloPrefix
	}

	return prefix + name
}
