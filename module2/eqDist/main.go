package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	white = iota
	black
)

// node - вершина графа
type node struct {
	color int
	depth int
	links []int
}

// clearGraph окрашивает все вершины в белый
func clearGraph(nodeList []*node) {
	for _, elem := range nodeList {
		elem.color = white
		elem.depth = 0
	}
}

// findEqDist находит все вершины, равноудаленные от опорных
func findEqDist(nodeList []*node, bearingList []int, isBearing []bool) []int {
	isEqDist := make([]bool, len(nodeList))
	bearDist := make([]int, len(nodeList))
	for _, rootIndex := range bearingList {
		clearGraph(nodeList)
		curBearNum := len(bearingList)
		nodeList[rootIndex].color = black
		queue := make([]int, 0)
		queue = append(queue, rootIndex)
		for len(queue) != 0 {
			curIndex := queue[0]
			queue = queue[1:]
			if isBearing[curIndex] {
				curBearNum--
			} else if bearDist[curIndex] == 0 {
				bearDist[curIndex] = nodeList[curIndex].depth
				isEqDist[curIndex] = true
			} else if bearDist[curIndex] != nodeList[curIndex].depth {
				isEqDist[curIndex] = false
			}
			for _, nextIndex := range nodeList[curIndex].links {
				if nodeList[nextIndex].color == white {
					nodeList[nextIndex].color = black
					nodeList[nextIndex].depth = nodeList[curIndex].depth + 1
					queue = append(queue, nextIndex)
				}
			}
		}
		if curBearNum != 0 {
			return nil
		}
	}
	result := make([]int, 0)
	for i, elem := range isEqDist {
		if elem {
			result = append(result, i)
		}
	}
	return result
}

func readSlice(scanner *bufio.Scanner) []int {
	scanner.Scan()
	result := make([]int, 0)
	for _, elem := range strings.Fields(scanner.Text()) {
		x, _ := strconv.Atoi(elem)
		result = append(result, x)
	}
	return result
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	input := strings.TrimSuffix(scanner.Text(), "\n")
	x, _ := strconv.Atoi(input)
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = writer.Flush()
	}()

	nodeNum := readInt(scanner)
	edgeNum := readInt(scanner)
	nodeList := make([]*node, 0)
	for i := 0; i < nodeNum; i++ {
		nodeList = append(nodeList, &node{white, 0, make([]int, 0)})
	}
	for i := 0; i < edgeNum; i++ {
		edge := readSlice(scanner)
		x, y := edge[0], edge[1]
		nodeList[x].links = append(nodeList[x].links, y)
		nodeList[y].links = append(nodeList[y].links, x)
	}

	bearingNum := readInt(scanner)
	isBearing := make([]bool, nodeNum)
	var bearingList []int
	if bearingNum > 0 {
		bearingList = readSlice(scanner)
	}
	for _, elem := range bearingList {
		isBearing[elem] = true
	}

	eqDistNodes := findEqDist(nodeList, bearingList, isBearing)
	for _, elem := range eqDistNodes {
		_, _ = writer.WriteString(strconv.Itoa(elem) + " ")
	}
	if len(eqDistNodes) == 0 {
		_, _ = writer.WriteString("-")
	}
	_, _ = writer.WriteString("\n")
}
