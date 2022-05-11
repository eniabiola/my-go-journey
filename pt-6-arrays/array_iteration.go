package main

import "fmt"

func main() {
	var arr = [5]float64{11, 24, 13, 15, 16}

	/*
		range helps to access the content of an array so the function below translates to
		loop through the array arr, and print out its its and corresponding value(content)
	*/
	for index, content := range arr {
		fmt.Println(index, " ", content)
	}
	//here the _ allows you to not use the index, if we declared index in the condition without using it,
	//Go will return an error as usual, it does not permit declaring a variable without using it.
	//This helps to manage memeory
	for _, content := range arr {
		fmt.Println(content)
	}
}
