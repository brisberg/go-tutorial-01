package main

import "fmt"

type triangle struct {
	base   int
	height int
}

type square struct {
	sideLength int
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		base:   10,
		height: 5,
	}

	s := square{
		sideLength: 6,
	}

	fmt.Println(printArea(t))
	fmt.Println(printArea(s))
}

func (t triangle) getArea() float64 {
	return 0.5 * float64(t.base*t.height)
}

func (s square) getArea() float64 {
	return float64(s.sideLength * s.sideLength)
}

func printArea(s shape) float64 {
	return s.getArea()
}
