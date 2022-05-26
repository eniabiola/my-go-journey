package main

import "fmt"

func main() {
	//A slice is a reference to an array, unlike when you copy an array, the copied array maintains its own
	//value but when you create a slice of an array, any modification made to that slice affects the original
	//array i.e slicedArrays are reference types of their original arrays
	arr1 := [...]float64{100, 30, 50.5, 70.9, 90, 100.6}
	fmt.Println("original array before slicing", arr1)
	slc := arr1[1:4]
	slc[0] = 2
	slc[1] = 40.9
	slc[2] = 90.8
	fmt.Println("sliced array", slc)
	fmt.Println("original array after slicing", arr1)
}
