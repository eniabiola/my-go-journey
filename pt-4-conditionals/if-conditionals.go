package main

import "fmt"

func main() {
	num1 := 50
	num2 := 100

	if num1 > num2 {
		fmt.Println(num1, "is greater than ", num2)
	} else if num1 == num2 {
		fmt.Println(num1, "is equal to ", num2)
	} else {
		fmt.Println(num1, "is less than ", num2)
	}

	//Conditional Variant
	if num3 := 4; num3%2 == 0 {
		fmt.Println("It is an even number")
	}
}
