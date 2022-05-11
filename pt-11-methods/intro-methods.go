package main

import "Tutorials/computer"

func main() {

	p1 := computer.ProcessingUnit{
		Price:        1000,
		Maker:        "IBM",
		MHz:          12000,
		Architecture: "x86",
	}

	p1.PrintSummary()
	p1.GiveDiscount(600)
	p1.PrintSummary()
}
