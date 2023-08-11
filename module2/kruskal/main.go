package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Лес непересекающихся множеств взят из презентации по АиСД
// Алгоритм Крускала взят из презентации по ДМ

// coordinate - координата аттракциона
type coordinate struct {
	x int
	y int
}

// edge - ребро графа
type edge struct {
	a     *set
	b     *set
	value float64
}

// set - дерево в лесу непересекающихся множеств
type set struct {
	coord  coordinate
	depth  int
	parent *set
}

// makeSet инициализирует дерево
func makeSet(c coordinate) *set {
	newSet := &set{
		coord: c,
		depth: 0,
	}
	newSet.parent = newSet
	return newSet
}

// findRoot находит корень дерева
func findRoot(x *set) *set {
	if x.parent == x {
		return x
	}
	return findRoot(x.parent)
}

// union объединяет деревья
func union(a, b *set) {
	rootA := findRoot(a)
	rootB := findRoot(b)
	if rootA.depth < rootB.depth {
		rootA.parent = rootB
	} else {
		rootB.parent = rootA
		if rootA.depth == rootB.depth && rootA != rootB {
			rootA.depth++
		}
	}
}

// calcDist считает расстояние между аттракционами
func calcDist(a, b *set) float64 {
	return math.Sqrt(math.Pow(float64(b.coord.x-a.coord.x), 2) +
		math.Pow(float64(b.coord.y-a.coord.y), 2))
}

// spanningTree формирует остовное дерево графа
func spanningTree(edges []*edge, nodeNum int) []*edge {
	result := make([]*edge, 0)
	for i := 0; i < len(edges) && len(result) < nodeNum-1; i++ {
		if findRoot(edges[i].a) != findRoot(edges[i].b) {
			result = append(result, edges[i])
			union(edges[i].a, edges[i].b)
		}
	}
	if len(result) != nodeNum-1 {
		panic("Несвязный граф!")
	}
	return result
}

// kruskal - алгоритм крускала
func kruskal(nodes []*set) []*edge {
	edges := make([]*edge, 0)
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			newEdge := &edge{
				a:     nodes[i],
				b:     nodes[j],
				value: calcDist(nodes[i], nodes[j]),
			}
			edges = append(edges, newEdge)
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].value < edges[j].value
	})
	return spanningTree(edges, len(nodes))
}

// totalLength считает общую длину дорожек
func totalLength(paths []*edge) float64 {
	var result float64
	for _, path := range paths {
		result += path.value
	}
	return result
}

// readCoordinate читает координату аттракциона из STDIN
func readCoordinate(reader *bufio.Reader) coordinate {
	input, _ := reader.ReadString('\n')
	trInput := strings.TrimSuffix(input, "\n")
	arr := strings.Split(trInput, " ")
	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])
	return coordinate{
		x: x,
		y: y,
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = writer.Flush()
	}()

	input, _ := reader.ReadString('\n')
	trInput := strings.TrimSuffix(input, "\n")
	n, _ := strconv.Atoi(trInput)
	nodes := make([]*set, 0)
	var minLength float64

	if n == 0 {
		// если аттракционов нет, то и дорожек нет
		minLength = 0
	} else {
		for i := 0; i < n; i++ {
			c := readCoordinate(reader)
			nodes = append(nodes, makeSet(c))
		}
		minPaths := kruskal(nodes)
		minLength = totalLength(minPaths)
	}

	_, _ = fmt.Fprintf(writer, "%.2f\n", minLength)
}
