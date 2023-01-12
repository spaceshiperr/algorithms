package search

import "algorithms/common"

func BreadthFirstSearch[T common.Numeric](graph common.Graph[T], root common.Vertex[T]) []T {
	queue := common.NewQueue()
	enqueueVertices(queue, root)

}

func enqueueVertices[T common.Numeric](queue common.Queue[T], vertex common.Vertex[T]) {
	for _, v := range vertex.Adjacent {
		queue.Enqueue(v.Key)
	}
}
