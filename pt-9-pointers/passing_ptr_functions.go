package main

import "fmt"

func main() {
	num1 := 100
	fmt.Println("num1 before modification", num1)
	modifyVariable(&num1)
	fmt.Println("num1 after modification by pointer", num1)
	modifyVariable2(num1)
	fmt.Println("num1 after modification by passing the variabke itself", num1)
}

func modifyVariable(ptr *int) {
	*ptr = 200
}

func modifyVariable2(num int) {
	num = 300
}
