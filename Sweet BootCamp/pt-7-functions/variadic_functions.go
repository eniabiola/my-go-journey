package main

import "fmt"

/*
THe use of variadic functions
*/

func main() {

	var summ float64
	summ = summation(10.0, 20.5, 20.5, 40.6, 90.4)
	fmt.Println(summ)

}

func summation(inputParam ...float64) (sum float64) {
	sum = 0
	for _, entry := range inputParam {
		sum += entry
	}
	return sum
}
