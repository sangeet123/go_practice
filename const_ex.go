package main

import (
	"fmt"
)

type Weekdays int

const (
	Sunday Weekdays = iota
	Monday
	Tuesday
	Wednesday
	Thrusday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Saturday)
}
