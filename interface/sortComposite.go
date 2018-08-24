package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	Name   string
	Age    int
	Salary float32
}

type byName []*Employee

func (es byName) Len() int {
	return len(es)
}

func (es byName) Less(i, j int) bool {
	return es[i].Name < es[j].Name
}

func (es byName) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func main() {
	employees := []*Employee{
		{"Ram Prashad Shrestha", 62, 250000.0},
		{"Devendra Prashad Dahal", 65, 280000.0},
		{"Ripendra Karki", 70, 260000.0},
		{"Mukunda Khatiwada", 65, 270000.0},
		{"Yogendra Lammichanne", 48, 280000.0},
		{"Pritam Bastola", 71, 900000.0},
	}

	sort.Sort(byName(employees))

	for _, v := range employees {
		fmt.Println(*v)
	}

	sort.Sort(sort.Reverse(byName(employees))) // returns new type reverse that is not exported and its Less method is called by indices reversed.

	fmt.Println("---------------reverse sorting---------------------")
	for _, v := range employees {
		fmt.Println(*v)
	}
}
