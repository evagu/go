package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectan struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectan) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	cir := Circle{x: 0, y: 0, radius: 5}
	rec := Rectan{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\nRec AreaL%f\n", getArea(cir), getArea(rec))
}
