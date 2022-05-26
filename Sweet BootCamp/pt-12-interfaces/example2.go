package main

import "fmt"

type human interface {
	sayHello() string
}

type man struct {
	greeting string
}

type woman struct {
	greeting string
}

type child struct {
	greeting string
}

func (m man) sayHello() string {
	return m.greeting
}

func (w woman) sayHello() string {
	return w.greeting
}

func (c child) sayHello() string {
	return c.greeting
}

func greetUser(h human) {
	fmt.Println(h.sayHello())
}

func main() {
	man1 := man{"Hello how you dey"}
	woman1 := woman{"Hello, ready for work?"}
	kid1 := child{"I'm hungry"}

	greetUser(man1)
	greetUser(woman1)
	greetUser(kid1)
}
