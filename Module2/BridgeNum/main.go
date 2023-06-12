package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	WHITE = iota
	BLACK
)

// Очередь
type queue struct {
	data    []int
	head    int
	tail    int
	counter int
}

// Вершина графа
type node struct {
	color       int
	inComponent bool
	parent      int
	links       []int
}

func initQueue(capacity int) *queue {
	return &queue{
		data:    make([]int, capacity),
		head:    0,
		tail:    0,
		counter: 0,
	}
}

func (q *queue) Enqueue(n int) {
	q.data[q.tail] = n
	q.tail++
	q.counter++
	if q.tail == len(q.data) {
		q.tail = 0
	}
}

func (q *queue) Dequeue() int {
	deqElement := q.data[q.head]
	q.head++
	q.counter--
	if q.head == len(q.data) {
		q.head = 0
	}
	return deqElement
}

func (q *queue) IsEmpty() bool {
	return q.counter == 0
}

// Добавляет индекс следующей вершины
func (n *node) Add(destIndex int) {
	n.links = append(n.links, destIndex)
	n.color = WHITE
}

// Читает ребра из Stdin
func readEdge(reader *bufio.Reader) (int, int) {
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	nodes := strings.Split(input, " ")
	nodeIndexA, _ := strconv.Atoi(nodes[0])
	nodeIndexB, _ := strconv.Atoi(nodes[1])
	return nodeIndexA, nodeIndexB
}

func readInt(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	inputInt, _ := strconv.Atoi(input)
	return inputInt
}

func visitNodeMarkComponents(startNodeIndex int, nodeList []*node) {
	nodeList[startNodeIndex].inComponent = true
	for _, currentNodeIndex := range nodeList[startNodeIndex].links {
		if !nodeList[currentNodeIndex].inComponent &&
			nodeList[currentNodeIndex].parent != startNodeIndex {
			visitNodeMarkComponents(currentNodeIndex, nodeList)
		}
	}
}

func countBridgesInComponent(nodeList []*node, nodeQueue *queue) int {
	bridges := -1
	for !nodeQueue.IsEmpty() {
		startNodeIndex := nodeQueue.Dequeue()
		if !nodeList[startNodeIndex].inComponent {
			visitNodeMarkComponents(startNodeIndex, nodeList)
			bridges++
		}
	}
	return bridges
}

func visitNodeMarkParents(startNodeIndex int, nodeList []*node, nodeQueue *queue) {
	nodeList[startNodeIndex].color = BLACK
	nodeQueue.Enqueue(startNodeIndex)
	for _, currentNodeIndex := range nodeList[startNodeIndex].links {
		if nodeList[currentNodeIndex].color == WHITE {
			nodeList[currentNodeIndex].parent = startNodeIndex
			visitNodeMarkParents(currentNodeIndex, nodeList, nodeQueue)
		}
	}
}

func calcBridges(nodeList []*node) int {
	var bridgeCounter int
	nodeQueue := initQueue(len(nodeList))

	for startNodeIndex := range nodeList {
		if nodeList[startNodeIndex].color == WHITE {
			visitNodeMarkParents(startNodeIndex, nodeList, nodeQueue)
			bridgeCounter += countBridgesInComponent(nodeList, nodeQueue)
		}
	}

	return bridgeCounter
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	nodeNum := readInt(reader)
	edgeNum := readInt(reader)

	// Список инцидентности
	nodeList := make([]*node, 0)
	for i := 0; i < nodeNum; i++ {
		nodeList = append(nodeList, &node{
			color:       WHITE,
			inComponent: false,
			parent:      -1,
			links:       make([]int, 0),
		})
	}

	for i := 0; i < edgeNum; i++ {
		nodeIndexA, nodeIndexB := readEdge(reader)
		nodeList[nodeIndexA].Add(nodeIndexB)
		nodeList[nodeIndexB].Add(nodeIndexA)
	}

	bridgeNum := calcBridges(nodeList)
	_, _ = writer.WriteString(strconv.Itoa(bridgeNum) + "\n")
}
