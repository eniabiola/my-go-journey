package main

import "fmt"

func main() {
	str := "Enitan"
	printBytes(str)
	str2 := "señor"
	printBytes(str2)
	printRune(str2)
}

func printBytes(name string) {
	//ASCII Characters
	for i := 0; i < len(name); i++ {
		fmt.Printf("The value %c and the hexadecimale value %x \n", name[i], name[i])
	}
}

func printRune(name string) {
	//UNICODE Characters
	r := []rune(name)
	for i := 0; i < len(r); i++ {
		fmt.Printf("The value %c and the hexadecimale value %x \n", r[i], r[i])
	}
}
