package main

import "fmt"

func main() {
	fmt.Println("Say Greeting to the user")

	chanVar := make(chan<- int)
	go greetingUser(chanVar)

	//x := <- chanVar // This will throw an error
}

func greetingUser(chanVar chan<- int) {
	fmt.Println("Say Greeting to user")

	chanVar <- 100
}