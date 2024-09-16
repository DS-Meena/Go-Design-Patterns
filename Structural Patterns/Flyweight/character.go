package main

import "fmt"

// Flyweight interface
type Character interface {
	Display(font string, size int) // set extrinsic data
}

// Concrete Flyweight
type ConcreteCharacter struct {
	symbol rune // intrinsic data only
}

func (c *ConcreteCharacter) Display(font string, size int) {
	// Display the character using the provided font and size
	fmt.Printf("Character: %c with Font: %s and Size: %d\n", c.symbol, font, size)
}

// Flyweight Factory
type CharacterFactory struct {
	characters map[rune]Character
}

func (f *CharacterFactory) GetCharacter(symbol rune) Character {
	if f.characters == nil {
		f.characters = make(map[rune]Character)
	}

	// already present in cache
	if char, ok := f.characters[symbol]; ok {
		return char
	}

	// Create new character
	char := &ConcreteCharacter{symbol}
	f.characters[symbol] = char
	return char
}
