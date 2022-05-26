package main

import (
	"fmt"
	"time"
)

func SayHello(data chan bool) {
	fmt.Println("Say hello sub-routine")
	time.Sleep(10 * time.Second)
	data <- true
}

func main() {
	//channels are pipes through which go routines communicate with one another
	fmt.Println("Opening statement")
	done := make(chan bool)

	go SayHello(done)

	x := <-done

	fmt.Println(x)
	fmt.Println("Main routine")
}
