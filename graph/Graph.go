package main

import (
	"container/list"
	"dataStructure/linear_list"
	"dataStructure/queue"
	"fmt"
	"sort"
	"strconv"
)

type Graph struct {
	vertexList []string // 顶点数组
	edges      [][]int  // 边对应的关系
	edgesNums  int      // 边数
	visited    []bool   // 是否访问
}

// 无向图的邻接表 表示
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
添加一条边,无向图
*/
func (g *Graph) AddEdges(v1 int, v2 int, weight int) {
	g.edges[v1][v2] = weight
	g.edges[v2][v1] = weight
	g.edgesNums++
}

/**
添加一条边,有向图
*/
func (g *Graph) AddDirectedEdges(v1 int, v2 int, weight int) {
	g.edges[v1][v2] = weight
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
邻接矩阵DFS非递归实现
*/
func (g *Graph) DfsOfNoRecursionWithMatrix(i int) {
	if i < 0 || i >= len(g.vertexList) {
		panic(fmt.Errorf("index out bound of %d\n", i))
	}
	stack := linear_list.NewLinkedStack()
	circleQueue := queue.NewCircleQueue(len(g.vertexList))
	fmt.Printf("%s ", g.vertexList[i])
	stack.Push(i)              // 存储已经访问用来回溯的顶点
	_ = circleQueue.EnQueue(i) // 存储已经访问的顶点
	g.visited[i] = true
	for !circleQueue.IsEmpty() {
		data, _ := circleQueue.DeQueue()
		k := data.(int)
		for j := 0; j < len(g.vertexList); j++ {
			if g.edges[k][j] != 0 {
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
					if g.edges[k][j] != 0 {
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
	stack := linear_list.NewLinkedStack() // 存储已经访问过的顶点,用来回溯
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
		}
		if !stack.IsEmpty() {
			node = stack.Pop().Data.(*EdgeNode)
			node = node.next
		}
	}
}

/**
从已知连通图构建最小生成树
Prim算法
 	g-已知图的顶点集合
	fromGraph-顶点之间的关系(边)
	t-最小生成树的顶点集合,初始为空
	minimumSpanningTree-最小生成树()
	g(uV,e) t(tV,e)
	1.初始化的图的顶点信息uV,tV初始为空
	2.从uV中选择一个顶点加入到tV中,并从uV中删除这个顶点
	3.寻找一条最小权重的边(ti,ui)->ui=(uV-tV),并将ui加入tV,边(ti,ui)加入t,从uV中删除ui
	4.重复3直到uV中的顶点为空,此时所有的边组成的就是最小生成树
普里姆算法
*/
func buildMinimumSpanningTree(g *Graph) [][]int {
	// 原图的顶点信息
	uV := make([]int, len(g.vertexList))
	for k := range g.vertexList {
		uV[k] = k
	}
	// 最小生成树的顶点信息
	var tV []int
	// 生成的目标最小生成树
	minimumSpanningTree := make([][]int, len(g.edges))
	for i := 0; i < len(g.edges); i++ {
		minimumSpanningTree[i] = make([]int, len(g.edges[i]))
	}
	tV = append(tV, uV[0])
	uV = append(uV[1:])
	l := len(uV)
	for i := 0; i < l; i++ {
		// 搜索权重最小的边(t,u),并将顶点u加入tv,从uV中删除顶点u
		minT, minU, weight := selectMinimumWeight(uV, tV, g)
		minimumSpanningTree[minT][minU] = weight
		tV = append(tV, minU)
		index := selectIndex(minU, uV)
		uV = append(uV[:index], uV[index+1:]...)
	}
	fmt.Printf("最小生成树的顶点信息%v\n", tV)
	return minimumSpanningTree
}

func selectIndex(minU int, uV []int) int {
	index := 0
	for k, v := range uV {
		if v == minU {
			index = k
			break
		}
	}
	return index
}

/**
寻找权重最小的边
 (t,u)
*/
func selectMinimumWeight(uV []int, tv []int, g *Graph) (minT int, minU int, weight int) {
	minWeight := 100
	for j := 0; j < len(tv); j++ {
		t := tv[j]
		for k := 0; k < len(uV); k++ {
			u := uV[k]
			if g.edges[t][u] != 0 && g.edges[t][u] < minWeight {
				minWeight = g.edges[t][u]
				minU = u
				minT = t
			}
		}
	}
	return minT, minU, minWeight
}

/** =================================
边集数组
*/
type Edge struct {
	begin  int // 边起点在数组中的下标
	end    int // 边终点在数组中的下标
	weight int // 边的权重
}

type EdgeArr []Edge

func newEdgeArr(g *Graph) EdgeArr {
	edges := make(EdgeArr, g.edgesNums)
	k := 0
	for i := 0; i < len(g.edges); i++ {
		for j := 0; j < len(g.edges[i]); j++ {
			if g.edges[i][j] != 0 {
				edges[k].begin = i
				edges[k].end = j
				edges[k].weight = g.edges[i][j]
				g.edges[i][j] = 0
				g.edges[j][i] = 0
				k++
			}
		}
	}
	return edges
}

// =========================

/**
克鲁斯卡尔求最小生成树
*/
func kruskal(g *Graph) {
	cr := make([]int, len(g.vertexList)) // 用来判断是否形成回路
	for i := 0; i < len(cr); i++ {
		cr[i] = -1
	}
	edges := newEdgeArr(g)
	sort.Sort(edges)
	sumW := 0
	for i := 0; i < len(edges); i++ {
		road1 := isCircleRoad(cr, edges[i].begin)
		road2 := isCircleRoad(cr, edges[i].end)
		if road1 != road2 { // 没有回路
			cr[road1] = road2
			fmt.Printf("边(%d,%d) weight=%d\n", edges[i].begin+1, edges[i].end+1, edges[i].weight)
			sumW += edges[i].weight
		}
	}
	fmt.Printf("sumW=%d\n", sumW)
}

/**
用来判断是否形成回路
*/
func isCircleRoad(cr []int, vertexIndex int) int {
	for cr[vertexIndex] != -1 {
		vertexIndex = cr[vertexIndex]
	}
	return vertexIndex
}

func (e EdgeArr) Less(i int, j int) bool {
	return e[i].weight < e[j].weight
}

func (e EdgeArr) Len() int {
	return len(e)
}

func (e EdgeArr) Swap(i int, j int) {
	e[i], e[j] = e[j], e[i]
}

/**
求源点v1到其他n-1个顶点的最短路径
迪杰斯特拉算法
*/
const maxDistance = 32657 // 用来表示如果2个顶点之间没有路径时候的权值
func dijkstra(graph Graph, v1 int) {
	distance := make([]int, len(graph.vertexList)) // 源点到顶点的路径长度
	path := make([]int, len(graph.vertexList))     // 顶点的前置端点 例p[1]=0 表示顶点1的前置端点是0 p[5]=4表示顶点5的前置端点为4  // p[5]=3表示顶点5的前置端点为3
	s := make([]bool, len(graph.vertexList))       // 存储顶点是否已经求得最短路径

	// 初始化d,s,p
	for i := 0; i < len(graph.edges); i++ {
		s[i] = false
		if graph.edges[v1][i] != 0 {
			path[i] = v1
			distance[i] = graph.edges[v1][i]
		} else {
			path[i] = -1
			distance[i] = maxDistance
		}
	}
	path[v1] = -1 // v1的前置顶点为-1
	s[v1] = true
	v, w := -1, -1
	for i := 1; i < len(graph.edges); i++ {
		minW := maxDistance
		for w = 0; w < len(graph.edges); w++ {
			if !s[w] && distance[w] < minW {
				v = w
				minW = distance[w]
			}
		}
		s[v] = true
		for w = 0; w < len(graph.edges); w++ {
			if !s[w] && distance[v]+graph.edges[v][w] < distance[w] {
				distance[w] = distance[v] + graph.edges[v][w]
				path[w] = v
			}
		}
	}
	fmt.Printf("%v\n", path)
}

/**
拓扑排序
*/
func topologicalSorting(g *Graph) {
	stack := linear_list.NewLinkedStack()       // 存储没有前驱的顶点
	precursor := make([]int, len(g.vertexList)) // 存储有前驱的顶点
	selectNoPrecursorVertex(g, stack, precursor)
	for !stack.IsEmpty() {
		i := stack.Pop().Data.(int)
		fmt.Printf("顶点:%s ", g.vertexList[i])
		precursor[i] = -1 // 删掉已经访问过的没有前驱的顶点
		for k := 0; k < len(g.edges[0]); k++ {
			if g.edges[i][k] == 1 {
				g.edges[i][k] = 0 // 删掉与该顶点相关联的边
				precursor[k] = 0  // 与之相关联的顶点k的前驱置空
			}
		}
		selectNoPrecursorVertex(g, stack, precursor)
	}
}

func selectNoPrecursorVertex(g *Graph, stack *linear_list.LinkedStack, precursor []int) {
	for i := 0; i < len(g.edges); i++ {
		for j := 0; j < len(g.edges[i]); j++ {
			if g.edges[i][j] == 1 && g.edges[j][i] == 0 {
				precursor[j] = i + 1
			}
		}
	}
	// fmt.Printf("precursor===%v \n", precursor)
	for i := 0; i < len(precursor); i++ {
		if precursor[i] == 0 {
			stack.Push(i)
			break
		}
	}
	// stack.Show()
}

func main() {
	// graph := NewGraph(6)
	// buildUnDirectedGraph(graph)
	// kruskal(graph)

	graph := NewGraph(9)
	buildDirectedGraph(graph)
	topologicalSorting(graph)

	// fmt.Println("深度优先===================")
	// graph.DfsOfNoRecursionWithMatrix(0)
	// fmt.Println()
	// ma := []bool{false, false, false, false, false, false, false, false}
	// graph.dfs(ma)

	// 最小生成树
	// tree := buildMinimumSpanningTree(graph)
	// fmt.Println()
	// for _, v := range tree {
	// 	fmt.Printf("%v\n", v)
	// }

	// fmt.Println("\n广度优先===================")
	// for i := 0; i < len(graph.vertexList); i++ {
	// 	graph.visited[i] = false
	// }
	// graph.Bfs()
	// ma = []bool{false, false, false, false, false, false, false, false}
	// fmt.Println()
	// graph.bfs(ma)

	// 邻接表表示无向图
	// graph := CreateUnDirectedGraph(4, 5)
	// graph.DfsOfNoRecursionWithAdj(0)
}

func buildUnDirectedGraph(graph *Graph) {

	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")
	graph.AddEdges(0, 1, 6)
	graph.AddEdges(0, 2, 1)
	graph.AddEdges(0, 3, 5)
	graph.AddEdges(1, 2, 5)
	graph.AddEdges(1, 4, 3)
	graph.AddEdges(2, 4, 6)
	graph.AddEdges(2, 3, 5)
	graph.AddEdges(2, 5, 4)
	graph.AddEdges(4, 5, 6)
	graph.AddEdges(3, 5, 2)
	graph.ShowGraph()
}

func buildDirectedGraph(graph *Graph) {

	graph.AddVertex("C1")
	graph.AddVertex("C2")
	graph.AddVertex("C3")
	graph.AddVertex("C4")
	graph.AddVertex("C5")
	graph.AddVertex("C6")
	graph.AddVertex("C7")
	graph.AddVertex("C8")
	graph.AddVertex("C9")
	graph.AddDirectedEdges(0, 3, 1)
	graph.AddDirectedEdges(0, 6, 1)
	graph.AddDirectedEdges(3, 7, 1)
	graph.AddDirectedEdges(3, 2, 1)
	graph.AddDirectedEdges(6, 2, 1)
	graph.AddDirectedEdges(2, 4, 1)
	graph.AddDirectedEdges(2, 7, 1)
	graph.AddDirectedEdges(1, 6, 1)
	graph.AddDirectedEdges(1, 8, 1)
	graph.AddDirectedEdges(8, 5, 1)
	graph.AddDirectedEdges(5, 4, 1)

	graph.ShowGraph()
}
