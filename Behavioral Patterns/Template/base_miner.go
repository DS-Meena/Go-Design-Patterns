package main

import "fmt"

// BaseMiner implements the template method
type BaseMiner struct {
	miner DataMiner
}

func (b *BaseMiner) mine() { // template method
	b.openFile()
	b.miner.extractData()
	b.miner.parseData()
	b.miner.analyzeData()
	b.sendReport()
	b.closeFile()
}

func (b *BaseMiner) openFile() {
	println("Opening file...")
}

func (b *BaseMiner) closeFile() {
	println("Closing file...")
}

func (b *BaseMiner) sendReport() {
	fmt.Println("Sending report...")
}
