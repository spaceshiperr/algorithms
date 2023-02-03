package common

import "errors"

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

func Equal(u Vertex, v Vertex) bool {
	return u.Key == v.Key
}

func NewGraph[T Numeric]() Graph {
	g := Graph{}
	g.Vertices = make([]*Vertex, 0)

	return g
}

func (g *Graph) AddVertex(key int) error {
	if containsVertex(g.Vertices, key) {
		return errors.New("has the key already")
	}

	g.Vertices = append(g.Vertices, &Vertex{
		Key: key,
	})

	return nil
}

func (g *Graph) AddEdge(to, from int) error {
	toV, fromV := getVertex(g.Vertices, to), getVertex(g.Vertices, from)

	if toV == nil || fromV == nil {
		return errors.New("vertex not found")
	}

	if containsVertex(fromV.Adjacent, toV.Key) {
		return errors.New("edge already exists")
	}

	fromV.Adjacent = append(fromV.Adjacent, toV)

	return nil
}

func getVertex(vertices []*Vertex, key int) *Vertex {
	for _, v := range vertices {
		if v.Key == key {
			return v
		}
	}

	return nil
}

func containsVertex(vertices []*Vertex, key int) bool {
	for _, v := range vertices {
		if v.Key == key {
			return true
		}
	}

	return false
}
