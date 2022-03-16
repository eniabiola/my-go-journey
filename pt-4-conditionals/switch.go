package main

import "fmt"

func main() {
	var option int
	option = 2
	switch option {
	case 1:
		fmt.Println("Go running")
	case 2:
		fmt.Println("Go jogging")
	case 3:
		fmt.Println("Take brisk walk")
	default:
		fmt.Println("Take a rest")

	}
}
