package main

import (
	"fmt"
	"math"
)

//interfaces help achieve polymorphism

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.perimeter())
	fmt.Println(g.area())
}

func main() {
	c := circle{
		radius: 49,
	}
	r := rectangle{
		width:  8,
		height: 12,
	}

	measure(c)
	fmt.Println()
	measure(r)
}
