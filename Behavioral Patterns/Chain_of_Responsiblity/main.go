package main

func main() {
	firstLevel := &FirstLevelSupport{}
	secondLevel := &SecondLevelSupport{}

	// set next handler for first
	firstLevel.SetNext(secondLevel)

	firstLevel.Handle("General Inquiry")
	firstLevel.Handle("Technical Issue")
	firstLevel.Handle("Billing Question") // none handled this query
}
