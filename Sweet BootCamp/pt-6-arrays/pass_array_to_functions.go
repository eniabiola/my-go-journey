package main

import "fmt"

func main() {
	arr1 := [3]int{
		10,
		20,
		11,
	}

	fmt.Println("Original Array ", arr1)

	fmt.Println("Modified Array ", modifiedArray(arr1))

}

func modifiedArray(newArray [3]int) (arr [3]int) {
	arr = newArray
	arr[1] = 100
	return arr
}