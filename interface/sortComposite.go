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

type CustomSort struct {
	employees []*Employee
	less      func(x, y *Employee) bool
}

func (cs CustomSort) Len() int {
	return len(cs.employees)
}

func (cs CustomSort) Less(i, j int) bool {
	return cs.less(cs.employees[i], cs.employees[j])
}

func (cs CustomSort) Swap(i, j int) {
	cs.employees[i], cs.employees[j] = cs.employees[j], cs.employees[i]
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
	// We have used pointer so that swap does not swap the entire field
	// it just swap the reference and that will make the code more efficient.
	es := []*Employee{
		{"Ram Prashad Shrestha", 62, 250000.0},
		{"Devendra Prashad Dahal", 65, 280000.0},
		{"Ripendra Karki", 70, 260000.0},
		{"Mukunda Khatiwada", 65, 270000.0},
		{"Yogendra Lammichanne", 48, 280000.0},
		{"Pritam Bastola", 71, 900000.0},
		{"Pritam Bastola", 71, 950000.0},
	}

	sort.Sort(byName(es))

	for _, v := range es {
		fmt.Println(*v)
	}

	sort.Sort(sort.Reverse(byName(es))) // returns new type reverse that is not exported and its Less method is called by indices reversed.

	fmt.Println("---------------reverse sorting---------------------")
	for _, v := range es {
		fmt.Println(*v)
	}

	sort.Sort(CustomSort{es, func(x, y *Employee) bool {
		if x.Name != y.Name {
			return x.Name < y.Name
		}
		if x.Age != y.Age {
			return x.Age < y.Age
		}
		if x.Salary != y.Salary {
			return x.Salary < y.Salary
		}
		return false
	}})

	fmt.Println("---------------CustomSort sorting---------------------")
	for _, v := range es {
		fmt.Println(*v)
	}

}
