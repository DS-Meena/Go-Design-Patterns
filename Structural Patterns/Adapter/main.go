package main

import "fmt"

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	fmt.Println(adapter.Request())
}
