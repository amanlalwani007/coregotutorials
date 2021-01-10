package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pi * c.radius
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}
func getArea(s shape) float64 {
	return s.area()
}

func main() {
	cir := circle{x: 0, y: 0, radius: 5}
	rec := Rectangle{width: 10, height: 5}
	fmt.Printf("Circle Area %f", getArea(cir))
	fmt.Printf("Rectanle Area %f", getArea(rec))

}
