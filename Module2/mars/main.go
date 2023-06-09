package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	UNDEFINED = iota
	GROUPED
	UNGROUPED
)

// Вершина графа
type node struct {
	linkedNodesIndices []int
	white              bool
	groupStatus        int
}

// Структура для подсчета оптимальной разницы
type linkedDelta struct {
	grouped   []int
	ungrouped []int
	delta     int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Алгоритм разбиения по подгруппам, проходит по компоненте
func analyze(currentNodeIndex int, nodes []*node, grouped, ungrouped []int, delta int) ([]int, []int, int, bool) {
	if nodes[currentNodeIndex].white {
		nodes[currentNodeIndex].white = false

		if nodes[currentNodeIndex].groupStatus == UNDEFINED {
			// Вершина, с которой начинаем обход компоненты, сразу в группе
			nodes[currentNodeIndex].groupStatus = GROUPED
			grouped = append(grouped, currentNodeIndex)
			delta++
		}

		for _, otherNodeIndex := range nodes[currentNodeIndex].linkedNodesIndices {
			if nodes[currentNodeIndex].groupStatus == nodes[otherNodeIndex].groupStatus {
				// Если соседние вершины обе в группе или не в группе
				return nil, nil, 0, false
			} else if nodes[currentNodeIndex].groupStatus == GROUPED &&
				nodes[otherNodeIndex].groupStatus == UNDEFINED {
				nodes[otherNodeIndex].groupStatus = UNGROUPED
				ungrouped = append(ungrouped, otherNodeIndex)
				delta--
			} else if nodes[currentNodeIndex].groupStatus == UNGROUPED &&
				nodes[otherNodeIndex].groupStatus == UNDEFINED {
				nodes[otherNodeIndex].groupStatus = GROUPED
				grouped = append(grouped, otherNodeIndex)
				delta++
			}

			if nodes[otherNodeIndex].white {
				// Если соседняя вершина еще не посещена, посещаем ее
				var haveSolution bool
				grouped, ungrouped, delta, haveSolution =
					analyze(otherNodeIndex, nodes, grouped, ungrouped, delta)

				if !haveSolution {
					return nil, nil, 0, false
				}
			}
		}
	}
	return grouped, ungrouped, delta, true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	n, _ := strconv.Atoi(input)

	nodes := make([]*node, 0)

	for i := 0; i < n; i++ {
		input, _ = reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		columns := strings.Split(input, " ")

		links := make([]int, 0)
		for j := 0; j < n; j++ {
			if strings.Compare(columns[j], "+") == 0 {
				links = append(links, j)
			}
		}

		newNode := node{links, true, UNDEFINED}
		nodes = append(nodes, &newNode)
	}

	diff := 0
	deltaList := make([]linkedDelta, 0)
	var grouped, ungrouped []int
	var delta int
	var haveSolution bool

	for i := range nodes {
		grouped, ungrouped, delta, haveSolution = analyze(i, nodes, make([]int, 0), make([]int, 0), 0)
		if !haveSolution {
			_, _ = writer.WriteString("No solution\n")
			break
		} else {
			diff += delta
			deltaList = append(deltaList, linkedDelta{grouped, ungrouped, delta})
		}
	}

	if haveSolution {
		results := make([]int, 0)

		// Сделаем разницу оптимальной (для меньшей лекс. разницы идем с конца)
		for i := len(deltaList) - 1; i >= 0; i-- {
			if abs(diff-2*deltaList[i].delta) < abs(diff) ||
				(abs(diff-2*deltaList[i].delta) == abs(diff)) &&
					(len(deltaList[i].grouped) > len(deltaList[i].ungrouped)) {
				diff -= 2 * deltaList[i].delta
				deltaList[i].grouped, deltaList[i].ungrouped = deltaList[i].ungrouped, deltaList[i].grouped
			}
		}

		for _, element := range deltaList {
			for _, index := range element.grouped {
				results = append(results, index+1)
			}
		}

		// Отсортируем наш ответ
		sort.Ints(results)

		for _, result := range results {
			_, _ = writer.WriteString(strconv.Itoa(result) + " ")
		}
	}
}
