package main

func main() {
	characters := []Character{&Soldier{}, &Tank{}, &Healer{}}

	healthBooster := &HealthBoostVisitor{}

	for _, character := range characters {
		character.Accept(healthBooster)
	}
}
