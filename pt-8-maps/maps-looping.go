package main

import "fmt"

func main() {
	fruitData := map[string]int{
		"apple":     20,
		"pineapple": 5,
		"mango":     21,
		"orange":    40,
	}
	fruitData["agbalumo"] = 45
	delete(fruitData, "mango")
	fmt.Println(fruitData["apple"])
	fmt.Println(len(fruitData))
	for fruit, quantity := range fruitData {
		fmt.Println("Fruit:", fruit, "Quantity", quantity)
	}
}
