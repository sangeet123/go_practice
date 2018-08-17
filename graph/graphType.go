package graph

import (
	"sort"
)

type Graph struct {
	nodes map[string]bool
	edges map[string]map[string]bool
}

func (g *Graph) getNodes() []string {
	nodes := []string{}

	for n := range g.nodes {
		nodes = append(nodes, n)
	}

	return nodes
}

func (g *Graph) AddNode(n string) {
	if g.nodes == nil {
		g.nodes = make(map[string]bool)
	}
	g.nodes[n] = true
}

func (g *Graph) AddEdges(from, to string) {
	// add nodes
	g.AddNode(from)
	g.AddNode(to)

	if g.edges == nil {
		g.edges = make(map[string]map[string]bool)
	}

	if g.edges[from] == nil {
		g.edges[from] = make(map[string]bool)
	}

	g.edges[from][to] = true
}

func (g *Graph) HasEdge(from, to string) bool {
	return g.edges[from][to]
}

func (g Graph) TopoSort() []string {

	seen := make(map[string]bool)

	order := make([]string, 0)

	var f func(string)

	f = func(n string) {
		seen[n] = true
		for neighbor := range g.edges[n] {
			if !seen[neighbor] {
				f(neighbor)
			}
		}
		order = append(order, n)
	}

	nodes := g.getNodes()

	sort.Strings(nodes)

	for _, n := range nodes {
		if !seen[n] {
			f(n)
		}
	}

	return order
}
