package main

import (
	"fmt"
)

const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "

	spanish = "Spanish"
	french  = "French"
)

func main() {
	fmt.Println(Hello("Mec", "Spanish"))
}

func Hello(str, language string) string {
	if str == "" {
		str = "World"
	}
	return getPrefix(language) + str
}

func getPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}
