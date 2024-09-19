package main

import "fmt"

// DataMiner defines the template method
type DataMiner interface {
	extractData()
	parseData()
	analyzeData()
}

// PDFMiner implements DataMiner
type PDFMiner struct{}

func (p *PDFMiner) extractData() {
	fmt.Println("Extracting data from PDF...")
}

func (p *PDFMiner) parseData() {
	fmt.Println("Parsing data from PDF...")
}

func (p *PDFMiner) analyzeData() {
	fmt.Println("Analyzing PDF Data...")
}

// CSVMiner implements DataMiner
type CSVMiner struct{}

func (c *CSVMiner) extractData() {
	fmt.Println("Extracting data from CSV...")
}

func (c *CSVMiner) parseData() {
	fmt.Println("Parsing data from CSV...")
}

func (c *CSVMiner) analyzeData() {
	fmt.Println("Analyzing CSV Data...")
}
