package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "My name is Enitan"

	fmt.Println(strings.Contains(str, "Enitan"))
	fmt.Println(strings.Count(str, "Enitan"))
	fmt.Println(strings.HasPrefix(str, "My"))

	newStr := strings.Join([]string{"John", "is", "a", "good", "boy"}, " ")
	fmt.Println(newStr)
}
