package main

import (
	"fmt"

	"github.com/pratice/ds-algo/12-graph/go/1-BFS-search/queue"
)

//adjancy matrix
//as vertexes start from 1, so we are using the array indexes
//from 1 onwards and leaving the 0 index, to ease our calculation
//1 means have edges
//0 means no edges

var queueArr *queue.CircularQueue = &queue.CircularQueue{
	Size:  10,
	Front: 0,
	Rear:  0,
	Arr:   make([]int, 10),
}

var Visited [50]int

var AdjMatrix [][]int = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 1, 1, 0, 0, 0},
	{0, 1, 0, 1, 0, 0, 0, 0},
	{0, 1, 1, 0, 1, 1, 0, 0},
	{0, 1, 0, 1, 0, 1, 0, 0},
	{0, 0, 0, 1, 1, 0, 1, 1},
	{0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0},
}

//we are starting the code execution from the indexe 1 onwards, to match with the vertex starting 1
//and not using the 0 index of the arrays

func BFS(vertex int) {
	fmt.Println(vertex)
	Visited[vertex] = 1
	queueArr.Enqueue(vertex)

	for !queueArr.IsEmpty() {
		u := queueArr.Dequeue()
		for v := 1; v < len(AdjMatrix); v++ {
			if AdjMatrix[u][v] == 1 && Visited[v] == 0 {
				fmt.Println(v) //visited on this line means printed
				Visited[v] = 1
				queueArr.Enqueue(v)
			}
		}
	}
}

func main() {
	// for i := 0; i <= 7; i++ {
	// 	for j := 0; j <= 7; j++ {
	// 		fmt.Printf("%d", AdjMatrix[i][j])
	// 	}
	// 	fmt.Println()
	// }
	BFS(1)

}
