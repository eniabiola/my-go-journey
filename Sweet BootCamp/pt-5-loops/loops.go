package main

import "fmt"

func main() {
	for i := 0; i <= 15; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 13 {
			break
		}
		fmt.Println(i)
	}

	//unlike what you know for while loop, in go while loop is a modified for loop, see example below

	sum := 1

	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
