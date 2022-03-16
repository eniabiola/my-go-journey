package main

import "fmt"

func main() {
	num1 := 100
	num2 := 50

	switch {
	case num1 == 100 && num2 < 100:
		fmt.Println("number 1 is equal to 100 and number 2 is less than 100")
	case num1 < 50:
		fmt.Println("Number 1 is less than 50")
	default:
		fmt.Println("Nothing special to see, just trying it conditional formatting")
	}

}
