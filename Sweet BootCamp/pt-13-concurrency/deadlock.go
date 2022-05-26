package main

import (
	"fmt"
	"time"
)

func main() {
	//deadlock happens when you're either sending and not receiving or you are receiving
	//and not sending via channels in go concurrency bidrectional subroutines
	chanVar := make(chan int)

	go greetUser()
	x := <-chanVar
	fmt.Println(x)

}

func greetUser() {
	fmt.Println("You're welcome")

	time.Sleep(5 * time.Second)
}
