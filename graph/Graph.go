package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	vertexList []string // 定点数组
	edges      [][]int  // 边对应的关系
	edgesNums  int      // 边数
}

func NewGraph(vertexNum int) *Graph {
	g := &Graph{
		vertexList: nil,
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
获得当前节点的第一个邻接节点
*/
func (g *Graph) getFirstNeighbor(currentIndex int) int {
	for i := 0; i < len(g.vertexList); i++ {
		if g.edges[currentIndex][i] > 0 {
			return i
		}
	}
	return -1
}

/**
根据 currentIndex 获得targetIndex后的一个邻接节点
*/
func (g *Graph) getNextNeighbor(currentIndex int, targetIndex int) int {
	for i := targetIndex + 1; i < len(g.vertexList); i++ {
		if g.edges[currentIndex][i] > 0 {
			return i
		}
	}
	return -1
}

/**
深度优先
*/
func (g *Graph) dfs1(startVertex int, ma []bool) {
	if startVertex > (len(g.vertexList) - 1) {
		panic("index out bounds of vertexList")
	}
	fmt.Printf("%s ", g.vertexList[startVertex])
	// 标记为已经访问
	ma[startVertex] = true

	neighbor := g.getFirstNeighbor(startVertex)
	// 存在一条边
	for neighbor != -1 {
		// 没有被访问过
		if !ma[neighbor] {
			g.dfs1(neighbor, ma)
		}
		neighbor = g.getNextNeighbor(startVertex, neighbor)
	}
}

func (g *Graph) dfs(ma []bool) {
	for i := 0; i < len(g.vertexList); i++ {
		if !ma[i] {
			g.dfs1(i, ma)
		}
	}
}

func (g *Graph) bfs1(startVertex int, ma []bool) {
	if startVertex > (len(g.vertexList) - 1) {
		panic("index out bounds of vertexList")
	}
	fmt.Printf("%s ", g.vertexList[startVertex])
	ma[startVertex] = true
	i := list.New()
	i.PushBack(startVertex)
	for i.Len() != 0 {
		front := i.Front()
		i.Remove(front)
		neighbor := g.getFirstNeighbor(front.Value.(int))
		for neighbor != -1 {
			if !ma[neighbor] {
				fmt.Printf("%s ", g.vertexList[neighbor])
				ma[neighbor] = true
				i.PushBack(neighbor)
			}
			neighbor = g.getNextNeighbor(front.Value.(int), neighbor)
		}
	}
}

/**
广度优先
*/
func (g *Graph) bfs(ma []bool) {
	for i := 0; i < len(g.vertexList); i++ {
		if !ma[i] {
			g.bfs1(i, ma)
		}
	}
}

func main() {
	graph := NewGraph(8)
	buildGraph(graph)

	ma := []bool{false, false, false, false, false, false, false, false}
	graph.dfs(ma)
	ma = []bool{false, false, false, false, false, false, false, false}

	fmt.Println()
	graph.bfs(ma)

}

func buildGraph(graph *Graph) {

	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")
	graph.AddVertex("7")
	graph.AddVertex("8")

	graph.AddEdges(0, 1, 1)
	graph.AddEdges(0, 2, 1)
	graph.AddEdges(1, 3, 1)
	graph.AddEdges(1, 4, 1)
	graph.AddEdges(3, 7, 1)
	graph.AddEdges(4, 7, 1)
	graph.AddEdges(2, 5, 1)
	graph.AddEdges(2, 6, 1)
	graph.AddEdges(5, 6, 1)

	graph.ShowGraph()
}
