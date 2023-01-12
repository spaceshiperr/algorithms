package common

import "errors"

type Graph[T Numeric] struct {
	Vertices []*Vertex[T]
}

type Vertex[T Numeric] struct {
	Key      T
	Adjacent []*Vertex[T]
}

func NewGraph[T Numeric]() Graph[T] {
	g := Graph[T]{}
	g.Vertices = make([]*Vertex[T], 0)

	return g
}

func (g *Graph[T]) AddVertex(key T) error {
	if containsVertex(g.Vertices, key) {
		return errors.New("has the key already")
	}

	g.Vertices = append(g.Vertices, &Vertex[T]{
		Key: key,
	})

	return nil
}

func (g *Graph[T]) AddEdge(to, from T) error {
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

func getVertex[T Numeric](vertices []*Vertex[T], key T) *Vertex[T] {
	for _, v := range vertices {
		if v.Key == key {
			return v
		}
	}

	return nil
}

func containsVertex[T Numeric](vertices []*Vertex[T], key T) bool {
	for _, v := range vertices {
		if v.Key == key {
			return true
		}
	}

	return false
}
