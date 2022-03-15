package main

import "fmt"

//There are 2 types of defined variables globally defined and custom scoped defined
//globally are defined outside func main but custom inside func main

var gpa int
var studentName string
var isGraduated bool

var (
	name string
	age  int
)

func main() {
	var (
		num1 int32
		num2 int64
	)
	num1 = 2
	num2 = 40
	//in go if you define a variable you hva eto use it else you will get an error
	lastname := "Awosanya" //type inference, the language automatically determines what type of variable it is,
	//the method used above is called shorthand
	hobby, phoneNumber := "Playing Games", "09098765434" //compounded variables on a single length and it only works if
	// one of the variable is a new one altho ugh they can be reassigned, but instead of using :=, we will only use =
	gpa = 3
	studentName = "Femi"
	isGraduated = true
	fmt.Println(gpa)
	fmt.Println(studentName)
	fmt.Println(isGraduated)
	fmt.Println("GPA ", gpa, " name ", studentName)
	fmt.Println(lastname)
	fmt.Println("my hobby is ", hobby, "and you can reach me on ", phoneNumber)

	//declaring and using complex variables
	a := 1 + 2i
	b := complex(2, 4)
	c := a + b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(int64(num1) + num2)

	//var

}
