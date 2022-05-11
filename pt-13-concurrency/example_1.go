package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is a main go routine")

	go sayHello()
	time.Sleep(1 * time.Second) // this function allows the system to slow down
	//without the above line the system will execute the main function and we will not see
	//the result of the go routine because the main function would have finished
	//execution  before the result of the independent go routine comes bavk

	fmt.Println("This statement is after the go routine")
}

func sayHello() {
	fmt.Println("This is a subroutine")
}
