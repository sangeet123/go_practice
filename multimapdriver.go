package main

import (
	"./multimap"
	"fmt"
)

func main() {
	m := multimap.MultiMap{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])
}
