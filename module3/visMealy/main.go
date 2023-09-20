package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printGraph(dmatrix [][]int, fmatrix [][]string, writer *bufio.Writer) {
	writer.WriteString("digraph {\n    rankdir = LR\n")
	for i := 0; i < len(dmatrix); i++ {
		for j := 0; j < len(dmatrix[i]); j++ {
			writer.WriteString(fmt.Sprintf("    %d -> %d ", i, dmatrix[i][j]))
			writer.WriteString(fmt.Sprintf("[label = \"%c(%s)\"]\n", rune('a'+j), fmatrix[i][j]))
		}
	}
	writer.WriteString("}\n")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	_, _ = strconv.Atoi(scanner.Text())

	dmatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		dmatrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			d, _ := strconv.Atoi(scanner.Text())
			dmatrix[i][j] = d
		}
	}
	fmatrix := make([][]string, n)
	for i := 0; i < n; i++ {
		fmatrix[i] = make([]string, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			fmatrix[i][j] = scanner.Text()
		}
	}

	printGraph(dmatrix, fmatrix, writer)
}
