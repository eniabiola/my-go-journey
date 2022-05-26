package main

import "fmt"

func main() {
	// Pointers are variables that contain memory addresses as their values
	// A variable name directly references a value
	// A pointer indirectly refers to a value
	// A pointer needs be declared before its use
	// Pointers are defined as var pointer *int
	// The & operator gives out the memory address of a variable, this is called referencing
	// The * operator gives out value of the variable, this is called de-referencing
	// The type of the pointer should be the same as the type of variable it is pointing to

	var firstname string
	var ptrfirstname *string

	firstname = "Enitan"
	ptrfirstname = &firstname

	fmt.Println(firstname)
	fmt.Println(ptrfirstname)  //address
	fmt.Println(*ptrfirstname) //variable in the address
}
