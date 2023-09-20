package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// dfs проходит автомат в глубину
func dfs(state int, counter *int, visited []bool, dmatrix [][]int,
	can map[int]int, ds []int) (map[int]int, []int) {
	visited[state] = true
	ds = append(ds, state)
	can[state] = *counter
	*counter++
	for _, next := range dmatrix[state] {
		if !visited[next] {
			can, ds = dfs(next, counter, visited, dmatrix, can, ds)
		}
	}
	return can, ds
}

// calcCanonic запускает обход автомата в глубину для вычисления канонической нумерации
func calcCanonic(startState int, dmatrix [][]int) (map[int]int, []int) {
	visited := make([]bool, len(dmatrix))
	canonic := make(map[int]int)
	deltas := make([]int, 0)
	return dfs(startState, new(int), visited, dmatrix, canonic, deltas)
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
	q0, _ := strconv.Atoi(scanner.Text())

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

	can, ds := calcCanonic(q0, dmatrix)

	writer.WriteString(fmt.Sprintf("%d\n%d\n%d\n", n, m, can[q0]))
	for _, state := range ds {
		for j := 0; j < m; j++ {
			writer.WriteString(fmt.Sprintf("%d ", can[dmatrix[state][j]]))
		}
		writer.WriteByte('\n')
	}
	for _, state := range ds {
		for j := 0; j < m; j++ {
			writer.WriteString(fmt.Sprintf("%s ", fmatrix[state][j]))
		}
		writer.WriteByte('\n')
	}
}
