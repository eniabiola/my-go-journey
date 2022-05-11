package main

import "fmt"

func main() {
	var twodimArray [2][3]float64
	twodimArray[0][1] = 45
	twodimArray[1][2] = 100

	//looping through a two-dimensional array (multidimensional array)
	for i, content := range twodimArray {
		for j, value := range content {
			fmt.Printf("i:%v j:%v value:%v \n", i, j, value)
		}
	}
}
