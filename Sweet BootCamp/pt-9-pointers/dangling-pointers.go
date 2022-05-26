package main

import "fmt"

func main() {
	//since a pointer is a variable with memory address, therefore a pointer not pointing to any memory yet is called
	//a dangling pointer
	var ptr *int

	if ptr == nil {
		fmt.Println("This is a dangling pointer")
	}

	var x int = 5
	ptr = &x

	fmt.Printf("The type of ptr is %T and its value is %v \n", ptr, ptr)
	fmt.Printf("The type of x is %T and its value is %v \n", x, x)

	//changing value of X via variable and via pointer
	x = 10
	fmt.Println("The new value of x is: ", x)
	*ptr = 20
	fmt.Println("The new value of x changed via pointer is ", x)

}
