package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {

	t := triangle{height: 3.5, base: 4.75}
	s := square{sideLength: 7.75}

	fmt.Println(s)
	fmt.Println(t)
	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height

}

func printArea(s shape) {
	fmt.Println(s.getArea())

}
