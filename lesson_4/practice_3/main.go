package main

import "fmt"

type shape interface {
	getArea() float32
}

type triangle struct {
	base   float32
	height float32
}

type square struct {
	sideLength float32
}

func main() {
	t := triangle{base: 20, height: 10}
	printShape(t)
	s := square{sideLength: 25}
	printShape(s)
}

func (t triangle) getArea() float32 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float32 {
	return s.sideLength * s.sideLength
}

func printShape(s shape) {
	fmt.Println(s.getArea())
}
