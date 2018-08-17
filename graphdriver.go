package main

import (
	"./graph"
	"fmt"
)

func main() {
	g := new(graph.Graph)

	g.AddEdges("algorithms", "data structures")
	g.AddEdges("calculus", "linear algebra")
	g.AddEdges("compilers", "data structures")
	g.AddEdges("compilers", "formal languages")
	g.AddEdges("compilers", "computer organization")
	g.AddEdges("data structures", "discrete math")
	g.AddEdges("databases", "data structures")
	g.AddEdges("discrete math", "intro to programming")
	g.AddEdges("formal languages", "discrete math")
	g.AddEdges("networks", "operating systems")
	g.AddEdges("operating systems", "data structures")
	g.AddEdges("operating systems", "computer organization")
	g.AddEdges("programming languages", "data structures")
	g.AddEdges("programming languages", "computer organization")

	fmt.Printf("%t\n", g.HasEdge("programming languages", "computer organization"))
	fmt.Printf("%t\n", g.HasEdge("compilers", "algorithms"))

	for i, course := range g.TopoSort() {
		fmt.Printf("%d: %s\n", i, course)
	}
}
