package main

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

type Circle struct {
	Point // represent center
	R     float64
}

func main() {
	c := Circle{
		Point: Point{X: 0, Y: 0},
		R:     3.14,
	}
	fmt.Println(c.Point)
}
