package main

import "fmt"

func main() {
	//anonymous functions are functions that do not have a name
	greet := func(username string) (greeting string) {
		return "Hello " + username + ", This is an anonymous function"
	}
	fmt.Println(greet("Enitan"))
	fmt.Println(greet("Abiola"))
}
