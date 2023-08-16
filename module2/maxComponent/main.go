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

// stack - стек
type stack struct {
	data []int
}

// node - вершина графа
type node struct {
	color int
	isRed bool
	links []int
}

// Push добавляет элемент в стек
func (s *stack) Push(n int) {
	s.data = append(s.data, n)
}

// Pop удаляет элемент из стека
func (s *stack) Pop() int {
	n := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return n
}

// IsEmpty проверяет стек на пустоту
func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}

// initStack инициализирует стек
func initStack() *stack {
	return &stack{data: make([]int, 0)}
}

// Add добавляет индекс следующей вершины
func (n *node) Add(destIndex int) {
	n.links = append(n.links, destIndex)
	n.color = WHITE
}

// Читает ребра из Stdin
func readEdge(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	nodes := strings.Fields(scanner.Text())
	nodeIndexA, _ := strconv.Atoi(nodes[0])
	nodeIndexB, _ := strconv.Atoi(nodes[1])
	return nodeIndexA, nodeIndexB
}

// clearGraph окрашивает все вершины в белый
func clearGraph(nodeList []*node) {
	for _, element := range nodeList {
		element.color = WHITE
	}
}

// analyzeComponent проходит по компоненте, возвращает кол-во вершин и ребер
func analyzeComponent(startNodeIndex int, nodeList []*node) (int, int) {
	nodeStack := initStack()
	var nodeCounter, edgeCounter int
	if nodeList[startNodeIndex].color == WHITE {
		nodeStack.Push(startNodeIndex)
		for !nodeStack.IsEmpty() {
			currentNodeIndex := nodeStack.Pop()
			if nodeList[currentNodeIndex].color == WHITE {
				nodeList[currentNodeIndex].color = BLACK
				nodeCounter++
				for _, otherNodeIndex := range nodeList[currentNodeIndex].links {
					if nodeList[otherNodeIndex].color == WHITE {
						nodeStack.Push(otherNodeIndex)
						edgeCounter++
					}
				}
			}
		}
	}
	return nodeCounter, edgeCounter
}

// findMaxComponent находит максимальную компоненту
func findMaxComponent(nodeList []*node) int {
	var maxNodes, maxEdges, maxComponentIndex int
	for startNodeIndex := range nodeList {
		nodes, edges := analyzeComponent(startNodeIndex, nodeList)
		if nodes > maxNodes {
			maxComponentIndex = startNodeIndex
			maxNodes = nodes
			maxEdges = edges
		} else if (nodes == maxNodes) && (edges > maxEdges) {
			maxComponentIndex = startNodeIndex
			maxEdges = edges
		}
	}
	// Покрасим все вершины обратно в белый
	clearGraph(nodeList)
	return maxComponentIndex
}

// printGraph выводит компоненту
func printGraph(nodeList []*node, maxComponentIndex int, writer *bufio.Writer) {
	nodeStack := initStack()
	_, _ = writer.WriteString("graph {\n")
	for startNodeIndex := range nodeList {
		if nodeList[startNodeIndex].color == WHITE {
			nodeStack.Push(startNodeIndex)
			for !nodeStack.IsEmpty() {
				currentNodeIndex := nodeStack.Pop()
				if nodeList[currentNodeIndex].color == WHITE {
					nodeList[currentNodeIndex].color = BLACK
					// выведем точку
					_, _ = writer.WriteString(strconv.Itoa(currentNodeIndex))
					if startNodeIndex == maxComponentIndex {
						_, _ = writer.WriteString(" [color=\"red\"]")
					}
					_, _ = writer.WriteString("\n")
					for _, otherNodeIndex := range nodeList[currentNodeIndex].links {
						if nodeList[otherNodeIndex].color == WHITE {
							nodeStack.Push(otherNodeIndex)
							// выведем ребро
							_, _ = writer.WriteString(strconv.Itoa(currentNodeIndex) + " -- " +
								strconv.Itoa(otherNodeIndex))
							if startNodeIndex == maxComponentIndex {
								_, _ = writer.WriteString(" [color=\"red\"]")
							}
							_, _ = writer.WriteString("\n")
						}
					}
				}
			}
		}
	}
	_, _ = writer.WriteString("}")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = writer.Flush()
	}()

	scanner.Scan()
	nodeNum, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	edgeNum, _ := strconv.Atoi(scanner.Text())

	// Список инцидентности
	nodeList := make([]*node, 0)
	for i := 0; i < nodeNum; i++ {
		nodeList = append(nodeList, &node{WHITE, false, make([]int, 0)})
	}

	for i := 0; i < edgeNum; i++ {
		nodeIndexA, nodeIndexB := readEdge(scanner)
		nodeList[nodeIndexA].Add(nodeIndexB)
		nodeList[nodeIndexB].Add(nodeIndexA)
	}

	// Найдем максимальную компоненту
	maxComponentIndex := findMaxComponent(nodeList)
	// Напечатем граф
	printGraph(nodeList, maxComponentIndex, writer)
}
