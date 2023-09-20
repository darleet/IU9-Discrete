package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// state - дерево в лесу непересекающихся множеств
type state struct {
	stNum  int
	depth  int
	parent *state
}

type machine struct {
	qNum      int
	xNum      int
	qStart    int
	dmatrix   [][]int
	fmatrix   [][]string
	visited   []bool
	canonical []int
}

// makeState инициализирует дерево
func makeState(s int) *state {
	newSet := &state{
		stNum: s,
	}
	newSet.parent = newSet
	return newSet
}

// findRoot находит корень дерева
func findRoot(x *state) *state {
	if x.parent == x {
		return x
	}
	return findRoot(x.parent)
}

// union объединяет деревья
func union(a, b *state) {
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

func split1(mealy *machine) ([]*state, int) {
	m := mealy.qNum
	q := make([]*state, 0)
	for i := 0; i < m; i++ {
		newSt := makeState(i)
		q = append(q, newSt)
	}
	for k, q1 := range q {
		for _, q2 := range q[k+1:] {
			if findRoot(q1) != findRoot(q2) {
				eq := true
				for i := 0; i < mealy.xNum; i++ {
					if mealy.fmatrix[q1.stNum][i] != mealy.fmatrix[q2.stNum][i] {
						eq = false
						break
					}
				}
				if eq {
					union(q1, q2)
					m--
				}
			}
		}
	}
	pi := make([]*state, 0)
	for _, qx := range q {
		pi = append(pi, findRoot(qx))
	}
	return pi, m
}

func split(mealy *machine, pi *[]*state) int {
	m := mealy.qNum
	q := make([]*state, 0)
	for i := 0; i < m; i++ {
		newSt := makeState(i)
		q = append(q, newSt)
	}
	for k, q1 := range q {
		for _, q2 := range q[k+1:] {
			if (*pi)[q1.stNum] == (*pi)[q2.stNum] && findRoot(q1) != findRoot(q2) {
				eq := true
				for i := 0; i < mealy.xNum; i++ {
					w1 := mealy.dmatrix[q1.stNum][i]
					w2 := mealy.dmatrix[q2.stNum][i]
					if (*pi)[w1] != (*pi)[w2] {
						eq = false
						break
					}
				}
				if eq {
					union(q1, q2)
					m--
				}
			}
		}
	}
	for _, qx := range q {
		(*pi)[qx.stNum] = findRoot(qx)
	}
	return m
}

func minimize(mealy *machine) *machine {
	var pi []*state
	var m1, m int
	pi, m1 = split1(mealy)
	for {
		m = split(mealy, &pi)
		if m == m1 {
			break
		}
		m1 = m
	}

	alreadyIn := make(map[int]bool)

	dmatrix1 := make([][]int, mealy.qNum)
	fmatrix1 := make([][]string, mealy.qNum)
	for i := range dmatrix1 {
		dmatrix1[i] = make([]int, mealy.xNum)
		fmatrix1[i] = make([]string, mealy.xNum)
	}
	mealy1 := &machine{
		qNum:      mealy.qNum,
		xNum:      mealy.xNum,
		qStart:    mealy.qStart,
		dmatrix:   dmatrix1,
		fmatrix:   fmatrix1,
		visited:   make([]bool, mealy.qNum),
		canonical: make([]int, mealy.qNum),
	}

	for i := 0; i < len(mealy.dmatrix); i++ {
		q1 := pi[i]
		if in, _ := alreadyIn[q1.stNum]; !in {
			alreadyIn[q1.stNum] = true
			for j := 0; j < len(mealy1.dmatrix[i]); j++ {
				mealy1.dmatrix[i][j] = pi[mealy.dmatrix[i][j]].stNum
				mealy1.fmatrix[i][j] = mealy.fmatrix[i][j]
			}
		}
	}
	return mealy1
}

// dfs проходит автомат в глубину
func dfs(mealy *machine, idx int, counter *int) {
	mealy.visited[idx] = true
	mealy.canonical[idx] = *counter
	*counter++
	for _, next := range mealy.dmatrix[idx] {
		if !mealy.visited[next] {
			dfs(mealy, next, counter)
		}
	}
}

// calcCanonic запускает обход автомата в глубину для вычисления канонической нумерации
func calcCanonic(mealy *machine) {
	counter := new(int)
	dfs(mealy, mealy.qStart, counter)
	mealy.qNum = *counter
	dmatrix := make([][]int, mealy.qNum)
	fmatrix := make([][]string, mealy.qNum)
	for i := range dmatrix {
		dmatrix[i] = make([]int, mealy.xNum)
		fmatrix[i] = make([]string, mealy.xNum)
	}
	for i := range mealy.canonical {
		if mealy.visited[i] {
			copy(dmatrix[mealy.canonical[i]], mealy.dmatrix[i])
			copy(fmatrix[mealy.canonical[i]], mealy.fmatrix[i])
			for j := range dmatrix[mealy.canonical[i]] {
				dmatrix[mealy.canonical[i]][j] = mealy.canonical[dmatrix[mealy.canonical[i]][j]]
			}
		}
	}
	mealy.dmatrix = make([][]int, len(dmatrix))
	mealy.fmatrix = make([][]string, len(fmatrix))
	mealy.qStart = 0
	copy(mealy.dmatrix, dmatrix)
	copy(mealy.fmatrix, fmatrix)
}

func printGraph(mealy *machine, writer *bufio.Writer) {
	writer.WriteString("digraph {\n    rankdir = LR\n")
	for i := 0; i < mealy.qNum; i++ {
		for j := 0; j < mealy.xNum; j++ {
			writer.WriteString(fmt.Sprintf("    %d -> %d ", i, mealy.dmatrix[i][j]))
			writer.WriteString(fmt.Sprintf("[label = \"%c(%s)\"]\n", rune('a'+j),
				mealy.fmatrix[i][j]))
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

	mealy := &machine{
		qNum:      n,
		xNum:      m,
		qStart:    q0,
		dmatrix:   dmatrix,
		fmatrix:   fmatrix,
		visited:   make([]bool, n),
		canonical: make([]int, n),
	}

	calcCanonic(mealy)
	mealy1 := minimize(mealy)
	calcCanonic(mealy1)
	printGraph(mealy1, writer)
}
