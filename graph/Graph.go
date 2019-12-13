package main

import (
	"container/list"
	"dataStructure/linear_list"
	"dataStructure/queue"
	"fmt"
	"strconv"
)

type Graph struct {
	vertexList []string // 顶点数组
	edges      [][]int  // 边对应的关系
	edgesNums  int      // 边数
	visited    []bool   // 是否访问
}

//
type EdgeNode struct {
	adjvex int       // 与当前节点相邻的顶点的在数组中的位置
	next   *EdgeNode // 下一个相邻的顶点
	weight int       // 权重
}

// 无向图的邻接表 表示
type VertexNode struct {
	vertex    string    // 顶点信息
	firstLink *EdgeNode // 指向邻接表的头结点
	visited   bool      // 是否被访问过
}

type UnDirectedGraph []VertexNode

func CreateUnDirectedGraph(vNum int, eNum int) UnDirectedGraph {
	g := make(UnDirectedGraph, vNum)
	for i := 0; i < vNum; i++ {
		g[i].vertex = "a" + strconv.Itoa(i)
		g[i].firstLink = nil
		g[i].visited = false
	}

	// 采用头插法创建无向图的邻接表
	vi, vj := 0, 0
	for k := 0; k < eNum; k++ {
		// 输入顶点(vi,vj)
		if _, err := fmt.Scanf("%d %d", &vi, &vj); err != nil {
			panic(err)
		}
		e := &EdgeNode{
			adjvex: vj,
			next:   nil,
		}
		e.next = g[vi].firstLink
		g[vi].firstLink = e

		e = &EdgeNode{
			adjvex: vi,
			next:   nil,
		}
		e.next = g[vj].firstLink
		g[vj].firstLink = e
	}
	return g
}

func NewGraph(vertexNum int) *Graph {
	g := &Graph{
		vertexList: nil,
		edges:      make([][]int, vertexNum),
		visited:    make([]bool, vertexNum),
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
	g.edges[v2][v1] = weight // 因为是无向图
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

// 邻接表的深度优先
func (g UnDirectedGraph) Dfs(i int) {
	if i < 0 || i >= len(g) {
		panic("index out bound ")
	}
	s := g[i].vertex
	fmt.Printf("%v\n", s)
	g[i].visited = true
	node := g[i].firstLink
	for node != nil {
		if !g[node.adjvex].visited {
			g.Dfs(node.adjvex)
		}
		node = node.next
	}
}

// 邻接表的广度优先
func (g UnDirectedGraph) Bfs() {
	circleQueue := queue.NewCircleQueue(len(g))
	i := 0
	node := g[i]
	fmt.Printf("%s ", node.vertex)
	g[i].visited = true
	_ = circleQueue.EnQueue(i)

	for !circleQueue.IsEmpty() {
		deQueue, _ := circleQueue.DeQueue()
		i = deQueue.(int)
		edgeNode := g[i].firstLink
		for edgeNode != nil {
			i = edgeNode.adjvex
			if !g[i].visited {
				node = g[i]
				fmt.Printf("%s ", node.vertex)
				g[i].visited = true
				_ = circleQueue.EnQueue(i)
			}
			edgeNode = edgeNode.next
		}
	}
}

// 邻接矩阵的深度优先
func (g *Graph) Dfs(i int) {
	if i < 0 || i >= len(g.vertexList) {
		panic(fmt.Errorf("out index of %d\n", i))
	}
	fmt.Printf("%s ", g.vertexList[i])
	g.visited[i] = true
	for j := 0; j < len(g.vertexList); j++ {
		if g.edges[i][j] == 1 && !g.visited[j] {
			g.Dfs(j)
		}
	}
}

// 邻接矩阵的广度优先
func (g *Graph) Bfs() {
	circleQueue := queue.NewCircleQueue(len(g.vertexList))
	i := 0
	g.visited[i] = true
	if err := circleQueue.EnQueue(i); err != nil {
		panic(err)
	}
	for !circleQueue.IsEmpty() {
		deQueue, err := circleQueue.DeQueue()
		if err != nil {
			panic(err)
		}
		i = deQueue.(int)
		fmt.Printf("%s ", g.vertexList[i])
		for j := 0; j < len(g.vertexList); j++ {
			if g.edges[i][j] == 1 && !g.visited[j] {
				g.visited[j] = true
				_ = circleQueue.EnQueue(j)
			}
		}
	}
}

/**
矩阵DFS非递归实现
*/
func (g *Graph) DfsOfNoRecursionWithMatrix(i int) {
	if i < 0 || i >= len(g.vertexList) {
		panic(fmt.Errorf("index out bound of %d\n", i))
	}
	stack := linear_list.NewLinkedStack()
	circleQueue := queue.NewCircleQueue(len(g.vertexList))
	fmt.Printf("%s ", g.vertexList[i])
	stack.Push(i)
	_ = circleQueue.EnQueue(i)
	g.visited[i] = true
	for !circleQueue.IsEmpty() {
		data, _ := circleQueue.DeQueue()
		k := data.(int)
		for j := 0; j < len(g.vertexList); j++ {
			if g.edges[k][j] == 1 {
				if !g.visited[j] {
					fmt.Printf("%s ", g.vertexList[j])
					stack.Push(j)
					_ = circleQueue.EnQueue(j)
					g.visited[j] = true
					break
				}
			}
		}

		if circleQueue.IsEmpty() && !stack.IsEmpty() {
			for !stack.IsEmpty() {
				k = stack.Pop().Data.(int)
				for j := 0; j < len(g.vertexList); j++ {
					if g.edges[k][j] == 1 {
						if !g.visited[j] {
							fmt.Printf("%s ", g.vertexList[j])
							_ = circleQueue.EnQueue(j)
							g.visited[j] = true
							break
						}
					}
				}
			}
		}
	}
}

/**
邻接表DFS非递归实现
*/
func (g UnDirectedGraph) DfsOfNoRecursionWithAdj(i int) {
	if i < 0 || i >= len(g) {
		panic(fmt.Errorf("index out bound of %d\n", i))
	}
	fmt.Printf("%s ", g[i].vertex)
	g[i].visited = true
	stack := linear_list.NewLinkedStack()
	node := g[i].firstLink
	stack.Push(node)
	for !stack.IsEmpty() || node != nil {
		for node != nil {
			if g[node.adjvex].visited {
				node = node.next
			} else {
				fmt.Printf("%s ", g[node.adjvex].vertex)
				g[node.adjvex].visited = true
				node = g[node.adjvex].firstLink
				stack.Push(node)
			}
			if !stack.IsEmpty() {
				node = stack.Pop().Data.(*EdgeNode)
				node = node.next
			}
		}
	}
}

func main() {
	// graph := NewGraph(8)
	// buildGraph(graph)
	// fmt.Println("深度优先===================")
	// graph.DfsOfNoRecursionWithMatrix(0)
	// fmt.Println()
	// ma := []bool{false, false, false, false, false, false, false, false}
	// graph.dfs(ma)
	// fmt.Println("\n广度优先===================")
	// for i := 0; i < len(graph.vertexList); i++ {
	// 	graph.visited[i] = false
	// }
	// graph.Bfs()
	// ma = []bool{false, false, false, false, false, false, false, false}
	// fmt.Println()
	// graph.bfs(ma)

	graph := CreateUnDirectedGraph(5, 6)
	graph.DfsOfNoRecursionWithAdj(0)
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
	// 1 2 4 8 5 3 6 7
	// 1 2 4 8 5 3 6 7

	graph.AddEdges(2, 3, 1)
	graph.AddEdges(4, 3, 1)
	graph.AddEdges(3, 5, 1)
	// 1 2 4 3 6 7 5 8
	// 1 2 4 3 6 7 5 8

	graph.ShowGraph()
}
