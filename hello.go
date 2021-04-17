package main

import ()

func main() {
}

const spanish string = "Spanish"
const french string = "French"
const english string = "English"
const italian string = "Italian"
const italianGreeting string = "Ciao, "
const englishGreeting string = "Hello, "
const spanishGreeting string = "Hola, "
const frenchGreeting string = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	languageMap := map[string]string{
		spanish: spanishGreeting,
		french:  frenchGreeting,
		english: englishGreeting,
		italian: italianGreeting,
	}

	return languageMap[language] + name
}
