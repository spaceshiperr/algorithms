package search

import (
	"algorithms/common"
)

func GetShortestPath(graph *common.DirectedGraph, from int) map[int]int {
	distances := make(map[int]int, 0)
	relaxed := make(map[int]bool)
	current := from

	for len(relaxed) != len(graph.Vertices) {
		for _, e := range graph.Edges[current] {
			if !relaxed[e.To] {
				if distances[e.From]+e.Weight < distances[e.To] || distances[e.To] == 0 {
					distances[e.To] = distances[e.From] + e.Weight
				}
			}
		}

		relaxed[current] = true

		current = findSmallest(graph, relaxed, distances)
	}

	return distances
}

func findSmallest(graph *common.DirectedGraph, relaxed map[int]bool, distances map[int]int) int {
	var smallest, next int

	for _, v := range graph.Vertices {
		if !relaxed[v.Id] {
			smallest = distances[v.Id]
			next = v.Id
		}
	}

	for _, v := range graph.Vertices {
		if !relaxed[v.Id] && v.Id != next {
			if distances[v.Id] != 0 && (distances[v.Id] < smallest && distances[v.Id] != 0 || smallest == 0) {
				smallest = distances[v.Id]
				next = v.Id
			}
		}
	}

	return next
}
