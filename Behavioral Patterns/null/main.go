package main

func main() {
	// with logging
	processWithLogger(&ConsoleLogger{})

	// without logging (using null object)
	processWithLogger(&NullLogger{})
}

func processWithLogger(logger Logger) {
	// some process
	logger.Log("Process completed")
}
