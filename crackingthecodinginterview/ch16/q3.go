package main

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

type Line struct {
	P1 Point
	P2 Point
}

func (l *Line) Slope() float64 {
	return (l.P2.Y - l.P1.Y) / (l.P2.X - l.P1.X)
}

func (l *Line) Intercept() float64 {
	return (l.P1.Y - l.Slope()*l.P1.X)
}

func IsPointInLs(p Point, l Line) bool {
	xWithingLineXs := (p.X >= l.P1.X && p.X <= l.P2.X) || (p.X >= l.P2.X && p.X <= l.P1.X)
	yWithingLineYs := (p.Y >= l.P1.Y && p.Y <= l.P2.Y) || (p.Y >= l.P2.Y && p.Y <= l.P1.Y)
	return xWithingLineXs && yWithingLineYs
}

func findPointWithinLines(l1 Line, l2 Line) *Point {
	if IsPointInLs(l1.P1, l2) {
		return &l1.P1
	}

	if IsPointInLs(l1.P2, l2) {
		return &l1.P2
	}

	if IsPointInLs(l2.P1, l1) {
		return &l2.P1
	}

	if IsPointInLs(l2.P2, l1) {
		return &l2.P2
	}

	return nil
}

// Does not consider the line with infinite slope i.e., vertical line
func GetIntersectionPoint(l1, l2 Line) *Point {

	s1 := l1.Slope()
	s2 := l2.Slope()

	i1 := l1.Intercept()
	i2 := l2.Intercept()

	// If both lines represent the same line
	if s1 == s2 && i1 == i2 {
		return findPointWithinLines(l1, l2)
	}

	// Two parallel lines cannot have intercept
	// to make thing readable second condition is kept
	if s1 == s2 && i1 != i2 {
		return nil
	}

	xIsect := (i1 - i2) / (s2 - s1)
	yIsect := (s1*xIsect + i1)
	pIsect := Point{xIsect, yIsect}

	if IsPointInLs(pIsect, l1) && IsPointInLs(pIsect, l2) {
		return &pIsect
	}

	return nil
}

func main() {

	// Overlapping lines
	l1 := Line{
		P1: Point{1, 1},
		P2: Point{5, 5},
	}

	l2 := Line{
		P1: Point{2, 2},
		P2: Point{6, 6},
	}

	fmt.Println(GetIntersectionPoint(l1, l2))
	fmt.Println(GetIntersectionPoint(l2, l1))

	// Non Overlapping lines
	l1 = Line{
		P1: Point{1, 1},
		P2: Point{5, 5},
	}

	l2 = Line{
		P1: Point{6, 6},
		P2: Point{10, 10},
	}

	fmt.Println(GetIntersectionPoint(l1, l2)) // expected nil
	fmt.Println(GetIntersectionPoint(l2, l1)) // expected nil

	// same slope line with different intercept
	l1 = Line{
		P1: Point{1, 1},
		P2: Point{5, 5},
	}

	l2 = Line{
		P1: Point{0, 3},
		P2: Point{4, 7},
	}

	fmt.Println(GetIntersectionPoint(l1, l2)) // expected nil

	// Intersecting line within lines point
	l1 = Line{
		P1: Point{1, 1},
		P2: Point{5, 5},
	}

	l2 = Line{
		P1: Point{0, 5},
		P2: Point{5, 0},
	}

	fmt.Println(GetIntersectionPoint(l1, l2)) // expected (2.5,2.5)

	// Intersecting line not within lines point
	l1 = Line{
		P1: Point{1, 1},
		P2: Point{3, 3},
	}

	l2 = Line{
		P1: Point{0, 10},
		P2: Point{10, 0},
	}

	fmt.Println(GetIntersectionPoint(l1, l2)) // expected nil

	// Vertical lines
	l1 = Line{
		P1: Point{1, 5},
		P2: Point{1, 10},
	}

	l2 = Line{
		P1: Point{1, 12},
		P2: Point{1, 8},
	}

	fmt.Println(GetIntersectionPoint(l1, l2)) // expected nil

}
