package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Функция поиска всех делителей числа
func findDividers(x int) (results []int) {
	for divider := 1; divider < int(math.Sqrt(float64(x)))+1; divider++ {
		if x%divider == 0 {
			results = append(results, divider)
			if x/divider != divider {
				results = append(results, x/divider)
			}
		}
	}
	sort.Ints(results)
	return results
}

// Функция поиска следующих вершин графа
func findNextNodes(dividers []int, nodeIndex int) (results []int) {
	for i := 0; i < nodeIndex; i++ {
		var dividerFits bool

		if dividers[nodeIndex]%dividers[i] == 0 {
			dividerFits = true
			for j := i + 1; j < nodeIndex; j++ {
				if dividers[j]%dividers[i] == 0 && dividers[nodeIndex]%dividers[j] == 0 {
					dividerFits = false
					break
				}
			}
		}

		if dividerFits {
			results = append(results, dividers[i])
		}
	}
	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	x, _ := strconv.Atoi(input)

	// найдем делители числа
	dividers := findDividers(x)

	_, _ = writer.WriteString("graph {\n")

	for i := len(dividers) - 1; i > 0; i-- {
		var results []int
		results = findNextNodes(dividers, i)
		for _, nextNode := range results {
			_, _ = writer.WriteString(strconv.Itoa(dividers[i]) + " -- " + strconv.Itoa(nextNode) + "\n")
		}
	}

	if len(dividers) == 1 {
		_, _ = writer.WriteString(input + "\n")
	}

	_, _ = writer.WriteString("}\n")
}
