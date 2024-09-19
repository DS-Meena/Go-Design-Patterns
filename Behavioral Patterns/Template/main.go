package main

import "fmt"

func main() {
	pdfMiner := &PDFMiner{}
	miner := BaseMiner{
		miner: pdfMiner,
	}

	fmt.Println("PDF Mining process")
	miner.mine()

	csvMiner := &CSVMiner{}
	miner = BaseMiner{
		miner: csvMiner,
	}

	fmt.Println("\nCSV Mining process")
	miner.mine()
}
