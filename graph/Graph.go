package main

import (
	"fmt"
)

type Graph struct {
	vertexList []string // 定点数组
	edges      [][]int  // 边对应的关系
	edgesNums  int      // 边数
}

func NewGraph(vertexNum int) *Graph {
	g := &Graph{
		vertexList: make([]string, vertexNum),
		edges:      make([][]int, vertexNum),
		edgesNums:  0,
	}
	for i := 0; i < len(g.edges); i++ {
		g.edges[i] = make([]int, vertexNum)
	}
	return g
}

/**
添加一条边
*/
func (g *Graph) AddEdges(v1 int, v2 int, weight int) {
	g.edges[v1][v2] = weight
	g.edges[v2][v1] = weight
	g.edgesNums++
}

/**
添加一个顶点
*/
func (g *Graph) AddVertex(vertex string) {
	g.vertexList = append(g.vertexList, vertex)
}

/**
打印
*/
func (g *Graph) ShowGraph() {
	for _, v := range g.edges {
		fmt.Printf("%v\n", v)
	}
}

/**
深度优先
*/
func (g *Graph) dfs(startVertex int) {
	if startVertex > (len(g.vertexList) - 1) {
		panic("index out bounds of vertexList")
	}
	// 用于表示节点是否被访问,1-表示未访问 0-表示已经访问
	ma := map[string]int{"A": 1, "B": 1, "C": 1, "D": 1, "E": 1}
	currentVertex := startVertex // 当前正在被访问的顶点

	s := g.vertexList[currentVertex]
	fmt.Printf("%s ", s)
	ma[s] = 0

bLabel:
	for index := 0; index < len(g.edges); index++ {

		for !allVisited(ma) {
			// 存在边
			visitable := (g.edges[currentVertex][index] == 1) && (g.edges[index][currentVertex] == 1)
			if visitable {
				// 没有被访问过
				if ma[g.vertexList[index]] == 1 {
					s := g.vertexList[currentVertex]
					fmt.Printf("%s ", s)
					ma[s] = 0
					currentVertex = index
					index = 0
				}
			} else {
				break bLabel
			}
		}
	}
}

func allVisited(ma map[string]int) bool {
	allVisited := true
	for _, v := range ma {
		if v == 1 {
			allVisited = false
			break
		}
	}
	return allVisited
}

func main() {
	graph := NewGraph(5)
	buildGraph(graph)

}

func buildGraph(graph *Graph) {

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")

	graph.AddEdges(0, 1, 1)
	graph.AddEdges(0, 2, 1)
	graph.AddEdges(1, 2, 1)
	graph.AddEdges(1, 3, 1)
	graph.AddEdges(1, 4, 1)
	graph.AddEdges(3, 1, 1)
	graph.AddEdges(4, 1, 1)

	graph.ShowGraph()
}
