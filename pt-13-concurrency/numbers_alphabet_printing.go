package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main sub routine ")

	go printNumbers()
	go printAlphabets()
	time.Sleep(6 * time.Second)
	fmt.Println("Main sub routine finished")
}

func printNumbers() {
	fmt.Println("This is a subroutine for printing numbers")

	for i := 1; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%v ", i)
	}
	fmt.Println("Print Numbers subroutine has finished")
}

func printAlphabets() {
	fmt.Println("This is a subroutine for printing alphabets")

	for i := 'a'; i < 'e'; i++ {
		time.Sleep(2 * time.Second)
		fmt.Printf("%c ", i)
	}
	fmt.Println("Print Alphabets subroutine has finished")
}
