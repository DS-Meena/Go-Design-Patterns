package main

// Client code
func main() {
	factory := &CharacterFactory{}

	text := "Hello, Flyweight! 🪶"
	for _, char := range text {
		character := factory.GetCharacter(char)
		character.Display("Arial", 12)
	}
}
