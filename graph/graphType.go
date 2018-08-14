package graph

type Graph struct {
	nodes map[string]bool
	edges map[string]map[string]bool
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
