package main

import "fmt"

func main() {
	//exportable functions are functions that are accessible to other files
	num1 := 50.0
	num2 := 10.0
	s, d := Calculator(num1, num2)
	fmt.Println("The sum is ", s)
	fmt.Println("The difference is ", d)
}

func Calculator(num1, num2 float64) (sum, difference float64) {
	difference = num1 - num2
	sum = num1 + num2
	return sum, difference
}
