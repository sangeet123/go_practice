package main

import (
	"./graph"
	"fmt"
)

func main() {
	g := new(graph.Graph)

	g.AddEdges("a", "b")
	g.AddEdges("b", "c")
	g.AddEdges("c", "d")
	g.AddEdges("a", "f")

	fmt.Printf("%t\n", g.HasEdge("a", "f"))
	fmt.Printf("%t\n", g.HasEdge("a", "c"))
}
