package main

import (
	"container/heap"
	"fmt"
	"math"
)

// node - вершина орграфа
type node struct {
	dist int
	x    int
	y    int
}

// queue - очередь с приоритетами, решил использовать
// встроенную реализацию кучи
type queue []*node

func (q *queue) Len() int {
	return len(*q)
}

func (q *queue) Less(i, j int) bool {
	return (*q)[i].dist < (*q)[j].dist
}

func (q *queue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *queue) Push(x any) {
	v := x.(*node)
	*q = append(*q, v)
}

func (q *queue) Pop() any {
	minEl := (*q)[0]
	(*q)[0] = (*q)[len(*q)-1]
	(*q)[len(*q)-1] = nil
	*q = (*q)[:len(*q)-1]
	return minEl
}

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

// dijkstra находит минимальное расстояние до конечной точки
func dijkstra(matrix [][]int) int {
	n := len(matrix)
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = math.MaxInt
		}
	}
	dist[0][0] = matrix[0][0]

	q := make(queue, 0)
	heap.Push(&q, &node{dist: matrix[0][0]})

	for !q.IsEmpty() {
		v := heap.Pop(&q).(*node)
		for _, offset := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newX, newY := v.x+offset[0], v.y+offset[1]
			if newX < n && newY < n && newX >= 0 && newY >= 0 &&
				v.dist+matrix[newX][newY] < dist[newX][newY] {
				u := node{
					dist: v.dist + matrix[newX][newY],
					x:    newX,
					y:    newY,
				}
				dist[newX][newY] = v.dist + matrix[newX][newY]
				heap.Push(&q, &u)
			}
		}
	}

	return dist[n-1][n-1]
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, err := fmt.Scan(&matrix[i][j])
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println(dijkstra(matrix))
}
