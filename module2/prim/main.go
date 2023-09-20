package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Алгоритм прима, очередь с приоритетами и heapify взяты из презентаций

// edge - ребро графа
type edge struct {
	to  int
	len int
}

// node - вершина графа
type node struct {
	index int
	key   int
	value *node
	links []*edge
}

// queue - очередь с приоритетами
type queue []*node

// heapify преобразует последовательность вершин в пирамиду
func (q *queue) heapify(i int) []*node {
	for {
		l := 2*i + 1
		r := l + 1
		root := i
		if l < len(*q) && (*q)[root].key > (*q)[l].key {
			root = l
		}
		if r < len(*q) && (*q)[root].key > (*q)[r].key {
			root = r
		}
		if i == root {
			break
		}
		(*q)[root], (*q)[i] = (*q)[i], (*q)[root]
		(*q)[root].index = root
		(*q)[i].index = i
	}
	return *q
}

// isEmpty проверяет очередь на пустоту
func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

// extractMin удаляет и возвращает минимальный элемент из очереди
func (q *queue) extractMin() *node {
	for i := len(*q)/2 - 1; i >= 0; i-- {
		q.heapify(i)
	}
	minEl := (*q)[0]
	minEl.index = -2
	(*q)[0] = (*q)[len(*q)-1]
	(*q)[len(*q)-1] = nil
	*q = (*q)[:len(*q)-1]
	q.heapify(0)
	return minEl
}

// decreaseKey уменьшает ключ вершины (в случае нахождения более короткого ребра)
func (q *queue) decreaseKey(n *node, newKey int) {
	n.key = newKey
}

// insert добавляет вершину в очередь
func (q *queue) insert(n *node) {
	*q = append(*q, n)
	n.index = len(*q) - 1
}

// prim реализует алгоритм Прима из презентации
func prim(graph []*node) int {
	q := make(queue, 0)
	var result int
	for v := graph[0]; ; {
		v.index = -2
		for _, link := range v.links {
			u := graph[link.to]
			if u.index == -1 {
				u.key = link.len
				u.value = v
				q.insert(u)
			} else if u.index != -2 && link.len < u.key {
				u.value = v
				q.decreaseKey(u, link.len)
			}
		}
		if q.isEmpty() {
			break
		}
		v = q.extractMin()
		result += v.key
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = writer.Flush()
	}()

	scanner.Scan()
	cabNum, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	roadNum, _ := strconv.Atoi(scanner.Text())

	graph := make([]*node, 0)
	for i := 0; i < cabNum; i++ {
		graph = append(graph, &node{
			index: -1,
			links: nil,
		})
	}
	for i := 0; i < roadNum; i++ {
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(input[0])
		b, _ := strconv.Atoi(input[1])
		length, _ := strconv.Atoi(input[2])
		graph[a].links = append(graph[a].links, &edge{to: b, len: length})
		graph[b].links = append(graph[b].links, &edge{to: a, len: length})
	}

	result := prim(graph)
	_, _ = writer.WriteString(strconv.Itoa(result) + "\n")
}
