package common

type DirectedGraph struct {
	Vertices []DirectedVertex
	Edges    map[int][]DirectedEdge
}

type DirectedEdge struct {
	Weight int
	From   int
	To     int
}

type DirectedVertex struct {
	Id int
}

func NewDirectedGraph(vertices []DirectedVertex, edges map[int][]DirectedEdge) *DirectedGraph {
	graph := &DirectedGraph{
		Vertices: vertices,
	}

	if edges != nil {
		graph.Edges = edges
	} else {
		graph.Edges = make(map[int][]DirectedEdge)
	}

	return graph
}

func (g *DirectedGraph) AddEdgeForDirectedGraph(from, to, weight int) {
	g.Edges[from] = append(g.Edges[from], DirectedEdge{
		Weight: weight,
		From:   from,
		To:     to,
	})
}

func (g *DirectedGraph) AddVertexForDirectedGraph(id int) {
	v := DirectedVertex{
		Id: id,
	}

	g.Vertices = append(g.Vertices, v)
	g.Edges[id] = make([]DirectedEdge, 0)
}
