package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	heigth float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.heigth * .5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
func main() {
	t := triangle{
		heigth: 12.5,
		base:   10.9}
	s := square{
		sideLength: 12.5}
	printArea(t)
	printArea(s)
}
