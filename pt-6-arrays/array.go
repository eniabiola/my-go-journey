package main

import "fmt"

func main() {
	var arr1 [4]float64
	arr1[0] = 23
	arr1[1] = 43
	fmt.Println(arr1)

	//shorter way of declaring arrays
	arr2 := [3]float64{
		11,
		22,
		33,
	}
	fmt.Println(arr2)
	//shorter way of declaring arrays when the size is not known at the point of declaration
	arr3 := [...]float64{
		11,
		22,
		33,
		44,
	}
	fmt.Println(arr3)

	//
	//In go arrays are value types and not reference types: i.e.  when you copy an array, and edit the copied array the original array remains intact
	arr4 := arr3
	arr4[0] = 23
	fmt.Println(arr4)
}
