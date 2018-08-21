package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func (p *Point) ScaleBy(s int) {
	p.x *= s
	p.y *= s
}

func main() {
	f := (*Point).ScaleBy

	p := Point{1, 2}

	f(&p, 3)

	fmt.Println(p)

	g := p.ScaleBy

	g(3)

	fmt.Println(p)
}
