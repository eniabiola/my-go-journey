package main

import (
	"Tutorials/geometry"
	"fmt"
)

func main() {
	length := 10.5
	var width float64 = 20

	area := geometry.Area(length, width)
	fmt.Println("Area is: ", area)

}
