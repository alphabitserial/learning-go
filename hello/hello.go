package main

import "fmt"

func main() {
	fmt.Println(Hello("Nadia", ""))
}

const (
	spanish             = "Spanish"
	french              = "French"
	mandarin            = "Mandarin"
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	mandarinHelloPrefix = "Ni hao, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case mandarin:
		prefix = mandarinHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
