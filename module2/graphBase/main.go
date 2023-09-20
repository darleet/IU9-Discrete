package main

import (
	"fmt"
	"sort"
	"strings"
)

// node представляет структуру вершины графа
type node struct {
	visitTime int
	lowTime   int
	comp      int
	links     []int
}

// component представляет структуру компоненты
type component struct {
	nodes []int
	links []int
}

// environment хранит данные о времени захода и номере компоненты
type environment struct {
	time  int
	count int
}

// visitNode посещает вершину (часть алгоритма Тарьяна из презентации)
func visitNode(graph, stack []*node, v *node, env *environment) []*node {
	v.visitTime = env.time
	v.lowTime = env.time
	env.time++
	stack = append(stack, v)
	for _, next := range v.links {
		u := graph[next]
		if u.visitTime == 0 {
			stack = visitNode(graph, stack, u, env)
		}
		if u.comp == 0 && v.lowTime > u.lowTime {
			v.lowTime = u.lowTime
		}
	}
	if v.visitTime == v.lowTime {
		for {
			u := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			u.comp = env.count
			if u == v {
				break
			}
		}
		env.count++
	}
	return stack
}

// tarjan находит компоненты сильной связности (из презентации)
func tarjan(graph []*node, env *environment) {
	stack := make([]*node, 0)
	for _, v := range graph {
		if v.visitTime == 0 {
			visitNode(graph, stack, v, env)
		}
	}
}

// buildCondensation строит конденсацию орграфа
func buildCondensation(graph []*node, env *environment) []*component {
	cond := make([]*component, 0)
	for i := 0; i < env.count-1; i++ {
		cond = append(cond, &component{
			nodes: make([]int, 0),
			links: make([]int, env.count-1),
		})
	}
	for i := range graph {
		v := graph[i]
		cond[v.comp-1].nodes = append(cond[v.comp-1].nodes, i)
		for _, j := range v.links {
			u := graph[j]
			// ставим -1 для дуги u -> v
			if u.comp != v.comp {
				cond[u.comp-1].links[v.comp-1] = -1
			}
		}
	}
	return cond
}

// findBase вычисляет базу заданного орграфа
func findBase(graph []*node) []int {
	env := &environment{
		time:  1,
		count: 1,
	}
	tarjan(graph, env)
	cond := buildCondensation(graph, env)
	base := make([]*component, 0)
	for _, v := range cond {
		// посчитаем полустепень захода
		var degSum int
		for _, deg := range v.links {
			degSum += deg
		}
		// если полустепень захода равна нулю, то это вершина базы
		if degSum == 0 {
			base = append(base, v)
		}
	}
	// в result запишем мин. вершины компонент из базы
	result := make([]int, 0)
	for _, comp := range base {
		result = append(result, comp.nodes[0])
	}
	sort.Ints(result)
	return result
}

func main() {
	var nodeNum, edgeNum int
	_, err := fmt.Scan(&nodeNum, &edgeNum)
	if err != nil {
		panic(err)
	}

	graph := make([]*node, 0)
	for i := 0; i < nodeNum; i++ {
		graph = append(graph, &node{links: make([]int, 0)})
	}

	for i := 0; i < edgeNum; i++ {
		var a, b int
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			panic(err)
		}
		graph[a].links = append(graph[a].links, b)
	}

	base := findBase(graph)
	fmt.Println(strings.Trim(fmt.Sprint(base), "[]"))
}
