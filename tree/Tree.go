package main

import (
	"dataStructure/linear_list"
	"dataStructure/queue"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type Stack struct {
	top  int
	data []interface{}
}

func NewStack() (stack *Stack) {
	stack = &Stack{
		top:  -1,
		data: nil,
	}
	return
}

func (s *Stack) Push(data interface{}) {
	s.top++
	s.data = append(s.data, data)
}

func (s *Stack) Pop() (data interface{}) {
	if s.top == -1 {
		return
	}
	data = s.data[s.top]
	s.top--
	return
}

func (s *Stack) Show() {
	i := s.top
	for i != -1 {
		fmt.Printf("%v ", s.data[i])
		i--
	}
	fmt.Println()
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

type TreeNode struct {
	data          int
	left          *TreeNode // 左子树
	right         *TreeNode // 右子树
	parent        *TreeNode // 父节点
	balanceFactor int       // 平衡因子 当前节点的左子树与右子树的高度之差
	nodeSize      int
}

/**
后序遍历
*/
func postOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}
	if tree.left != nil {
		postOrderTraversal(tree.left)
	}
	if tree.right != nil {
		postOrderTraversal(tree.right)
	}
	fmt.Printf("node=%v->bf=%d ", tree.data, tree.balanceFactor)
}

/**
先序遍历
*/
func preOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Printf("node=%v <-> bf=%d childNodeSize=%d\n ", tree.data, tree.balanceFactor, tree.nodeSize)
	if tree.left != nil {
		preOrderTraversal(tree.left)
	}
	if tree.right != nil {
		preOrderTraversal(tree.right)
	}
}

/**
中序遍历
*/
func middleOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}
	middleOrderTraversal(tree.left)
	fmt.Printf("node=%v <-> bf=%d, ", tree.data, tree.balanceFactor)
	middleOrderTraversal(tree.right)
}

/**
层级遍历
*/
func tierOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}
	circleQueue := queue.NewCircleQueue(20)
	_ = circleQueue.EnQueue(tree)
	for !circleQueue.IsEmpty() {
		deQueue, _ := circleQueue.DeQueue()
		node := deQueue.(*TreeNode)
		fmt.Printf("node=%v <-> bf=%d, ", tree.data, tree.balanceFactor)
		if node.left != nil {
			left := node.left
			_ = circleQueue.EnQueue(left)
		}

		if node.right != nil {
			right := node.right
			_ = circleQueue.EnQueue(right)
		}
	}
}

const n = 100
const m = 2*n - 1

type HNode struct {
	weight int
	lChild int
	rChild int
	parent int
}
type HuffmanTree [m]*HNode

// HuffmanTree构造步骤
// 1.根据与n个权值(w1,w2...)对应的n个节点构成的只有根节点的二叉树集合F(t1,t2,...)
// 2.创建一个只有根节点的新树,从F中选取根节点权值最小的2棵树作为新树的左右孩子,并将这2个节点的权值之和作为新树的权值
// 3.将新树加入F,从F中删除上面选取的2棵树
// 4.重复23直到F中只有一棵树为止,这棵树就是Huffman树
// 总共需合并n-1次,且产生n-1个新的节点
func NewHuffmanTree() (huffmanTree HuffmanTree) {

	huffmanTree = HuffmanTree{}
	for i := 0; i < len(huffmanTree); i++ {
		huffmanTree[i] = &HNode{
			weight: 0,
			lChild: 0,
			rChild: 0,
			parent: 0,
		}
	}
	initWeight(&huffmanTree)
	return
}
func (tree *HuffmanTree) Len() int {
	return len(tree)
}

func (tree *HuffmanTree) Less(i int, j int) bool {
	if tree[j].weight < tree[i].weight {
		return true
	}
	return false
}

func (tree *HuffmanTree) Swap(i int, j int) {
	tree[i], tree[j] = tree[j], tree[i]
}
func initWeight(tree *HuffmanTree) {
	if tree == nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	// 初始化节点的权值
	for i := 1; i <= n; i++ {
		tree[i].weight = rand.Intn(100)
	}
	// 以及构造Huffman 需要将原来的n个节点进行合并,需要n-1次合并,产生n-1个新节点
	for i := n + 1; i <= m; i++ {
		// 从n个只有根的二叉树节点中选择2个权值最小的
		minimum, nextMinimum := select2MinWeight(tree, n)
		// 将新节点设置为上面选取的2个节点的父节点
		tree[minimum].parent = i
		tree[nextMinimum].parent = i
		// 将上面选取的2个节点分别设置为新节点的左孩子以及右孩子
		tree[i].lChild = minimum
		tree[i].rChild = nextMinimum
		// 将上面选取的2个节点分别的权值之和作为新节点的权值
		tree[i].weight = tree[minimum].weight + tree[nextMinimum].weight
	}
	// sort.Sort(tree)
}

// 选择最小的以及第二小的2个节点
func select2MinWeight(tree *HuffmanTree, k int) (minimum int, nextMinimum int) {
	min := tree[0].weight
	for i := 1; i < k; i++ {
		if tree[i].weight < min && tree[i].parent == 0 {
			min = tree[i].weight
			minimum = i
		}
	}

	min = tree[0].weight
	for i := 0; i < k && i != minimum && tree[i].parent == 0; i++ {
		if tree[i].weight < min {
			min = tree[i].weight
			nextMinimum = i
		}
	}
	return
}
func testHuffmanTree() {
	huffmanTree := NewHuffmanTree()
	for k, v := range huffmanTree {
		fmt.Printf("k=%d v=%v\n", k, v)
	}
}

func isExpress(c string) (ok bool) {

	switch c {
	case "+", "-", "*", "/":
		return true
	}
	return false
}

func isNumber(c string) (matched bool) {
	matched, err := regexp.Match("^\\d+$", []byte(c))
	if err != nil {
		panic(err)
	}
	return matched
}

// 搜索二叉树
func (root *TreeNode) add(data int) *TreeNode {
	if node, cn := root.contains(data); cn {
		return node
	}
	if root == nil {
		return &TreeNode{
			data:          data,
			left:          nil,
			right:         nil,
			balanceFactor: 0,
			nodeSize:      1,
		}
	}
	if data < root.data {
		root.left = root.left.add(data)
		root.left.parent = root
	} else {
		root.right = root.right.add(data)
		root.right.parent = root
	}
	root.nodeSize++
	return root
}

func (root *TreeNode) contains(data int) (node *TreeNode, contain bool) {

	if root == nil {
		return nil, false
	}
	if data == root.data {
		return root, true
	} else if data > root.data {
		return root.right.contains(data)
	} else if data < root.data {
		return root.left.contains(data)
	}
	return nil, false
}

/**
删除节点
*/
func (*TreeNode) remove(tree *TreeNode, target int) *TreeNode {

	if tree == nil { // 空树
		return nil
	}
	if target == tree.data {
		if tree.left != nil && tree.right != nil {
			tree.data = tree.right.findMin().data
			fmt.Printf("min == %d \n", tree.data)
			tree.right = tree.remove(tree.right, tree.data)
		} else if tree.left == nil && tree.right == nil {
			tree = nil
		} else {
			temp := tree.parent
			if tree.left != nil {
				tree = tree.left
			} else {
				tree = tree.right
			}
			tree.parent = temp
		}
	} else if target > tree.data {
		tree.right = tree.remove(tree.right, target)
	} else if target < tree.data {
		tree.left = tree.left.remove(tree.left, target)
	}
	return tree
}

func (root *TreeNode) findMin() (minNode *TreeNode) {
	if root == nil {
		return nil
	}
	if root.left != nil {
		return root.left.findMin()
	}
	return root
}

func (root *TreeNode) findMax() (maxNode *TreeNode) {
	if root == nil {
		return nil
	}
	if root.right != nil {
		return root.right.findMax()
	}
	return root
}

/**
获取树的高度
*/
func (root *TreeNode) height() int {
	if root == nil {
		return 0
	}
	return int(1 + math.Max(float64(root.left.height()), float64(root.right.height())))
}

/**
判断是否是满二叉树
*/
func (root *TreeNode) isFullBinaryTree(tree *TreeNode) (bool, *TreeNode) {
	k := root.height()
	// 具有n的节点的完全二叉树的深度为Log(n)+1或者Log(n+1)  不是完全二叉树不能用来判断是否是满二叉树
	// if int(math.Log2(float64(root.nodeSize+1))) != k && int(math.Log2(float64(root.nodeSize))+1) != k {
	// 	return false
	// }
	// 深度为K的二叉树最多有2的k次方-1个节点
	// 节点数=2的k次方-1 是满二叉树
	if root.nodeSize == 1 {
		return false, root
	}
	if root.nodeSize == int(math.Pow(2, float64(k))-1) {
		return true, nil
	}
	if root.left != nil && root.right == nil {
		return false, root
	} else {
		f, node := root.left.isFullBinaryTree(tree)
		if f {
			return root.right.isFullBinaryTree(tree)
		} else {
			return f, node
		}
	}
}

/**
平衡二叉搜索树
*/
func (root *TreeNode) avlTreeAdd(data int) *TreeNode {
	_ = root.add(data)
	root.calTreeMinimumBalanceFactor() // 计算平衡因子
	minFactorNode := root.selectMinimumNonBalanceTree(data)
	if minFactorNode.balanceFactor > 1 { // 右旋
		return minFactorNode.rightRotate(minFactorNode.left) // 右旋的时候选取当前旋转节点的左孩子作为支点
	} else if minFactorNode.balanceFactor < -1 { // 左旋
		return minFactorNode.leftRotate(minFactorNode.right) // 左旋的时候选取当前旋转节点的右孩子作为支点
	} else {
		//
	}
	return root
}

/**
右旋
*/
func (root *TreeNode) rightRotate(childNode *TreeNode) *TreeNode {
	fmt.Printf("%s\n", "rightRotate")
	if childNode.balanceFactor != 0 { // 当前节点的平衡因子和子树的平衡因子的数学符号不一致,需要重新构造节点
		if pm := root.comPm(childNode); !pm {
			var temp *TreeNode
			if childNode.balanceFactor < 0 { // 说明右子树比左子树高,右子树一定不为空
				temp = childNode.singleLeftRotate(childNode)
			} else { // 说明左子树比右子树高,左子树一定不为空
				temp = childNode.singleRightRotate(childNode)
			}
			root.calTreeMinimumBalanceFactor()
			childNode = temp
		}
	}
	childNode.parent = root.parent
	root.left = childNode.right
	childNode.right = root
	if root.parent != nil {
		// root.parent.left = childNode
		if root.data > root.parent.data {
			root.parent.right = childNode
		} else {
			root.parent.left = childNode
		}
		p := root.parent
		for p != nil && p.parent != nil {
			p = p.parent
		}
		root.parent = childNode
		return p
	}
	root.parent = childNode
	return childNode
}

/**
左旋
*/
func (root *TreeNode) leftRotate(childNode *TreeNode) *TreeNode {
	fmt.Printf("%s\n", "leftRotate")
	// 当前需要旋转的节点的平衡因子与旋转支点的平衡因子的数学符号必须一致 (+ -一致)
	if childNode.balanceFactor != 0 {
		if pm := root.comPm(childNode); !pm {
			var temp *TreeNode
			fmt.Printf("%s\n", "leftRotate  comPm=============")
			if childNode.balanceFactor < 0 { // 右子树比左子树高 右子树一定不为空
				temp = childNode.singleLeftRotate(childNode)
			} else { // 说明左子树比右子树高,左子树一定不为空
				temp = childNode.singleRightRotate(childNode)
			}
			root.calTreeMinimumBalanceFactor() // 重新计算平衡因子
			childNode = temp                   // 因为重新调整了结构 child需要重新改变
		}
	}
	childNode.parent = root.parent
	childLeft := childNode.left
	if root.parent != nil {
		if root.data > root.parent.data {
			root.parent.right = childNode
		} else {
			root.parent.left = childNode
		}
		childNode.left = root
		root.right = childLeft
		root.parent = childNode
		p := root.parent
		for p != nil && p.parent != nil {
			p = p.parent
		}
		return p
	} // 如果是根节点
	root.right = childLeft
	root.parent = childNode
	childNode.left = root
	return childNode
}

/**
单左旋 用于平衡因子的符号不同的时候
*/
func (*TreeNode) singleLeftRotate(childNode *TreeNode) *TreeNode {
	temp := childNode.right
	node := temp.left
	temp.parent = childNode.parent
	childNode.parent.left = temp
	temp.left = childNode
	childNode.parent = temp
	childNode.right = node
	return temp
}

/**
单右旋 用于平衡因子的符号不同的时候
*/
func (*TreeNode) singleRightRotate(childNode *TreeNode) *TreeNode {
	temp := childNode.left
	node := temp.right
	temp.parent = childNode.parent
	childNode.parent.right = temp
	temp.right = childNode
	childNode.parent = temp
	childNode.left = node
	return temp
}

/**
比较当前节点的平衡因子和子树的平衡因子的符号是否相同
*/
func (root *TreeNode) comPm(childNode *TreeNode) bool {
	return balanceRegex("^-\\d", strconv.Itoa(root.balanceFactor), strconv.Itoa(childNode.balanceFactor))
}

func balanceRegex(partner string, target1 string, target2 string) bool {
	matched1, err := regexp.Match(partner, []byte(target1))
	if err != nil {
		panic(err)
	}
	matched2, err := regexp.Match(partner, []byte(target2))
	if err != nil {
		panic(err)
	}
	return matched1 == matched2
}

/**
计算每个节点的平衡因子(左子树与右子树高度之差 -1 0 1)
*/
func (root *TreeNode) calTreeMinimumBalanceFactor() {
	if root == nil {
		return
	}
	root.balanceFactor = root.left.height() - root.right.height()
	if root.right != nil {
		root.right.calTreeMinimumBalanceFactor()
	}
	if root.left != nil {
		root.left.calTreeMinimumBalanceFactor()
	}
}

/**
寻找最小不平衡子树
*/
func (root *TreeNode) selectMinimumNonBalanceTree(data int) *TreeNode {
	node := root
	stack := linear_list.NewLinkedStack()
	for math.Abs(float64(node.balanceFactor)) >= 1 {
		if math.Abs(float64(node.balanceFactor)) == 2 {
			stack.Push(node)
		}
		if data > node.data && node.right != nil {
			node = node.right
		} else if data < node.data && node.left != nil {
			node = node.left
		}
	}
	if !stack.IsEmpty() {
		return stack.Pop().Data.(*TreeNode)
	}
	return &TreeNode{
		data:          0,
		left:          nil,
		right:         nil,
		parent:        nil,
		balanceFactor: 0,
	}
}

// 树的孩子链表表示法
type ChildNode struct {
	child int        // 在数组中的下标
	next  *ChildNode // 下一个
}

type PanNode struct {
	data       int        // 数据域
	firstChild *ChildNode // 第一个孩子
}

type CTree struct {
	nodes     []PanNode
	nodeNum   int // 树中节点的数量
	rootIndex int // 树中根节点在数组中的位置
}

// 树的孩子链表表示法

func NewCTree(initCap int) *CTree {
	return &CTree{
		nodes:     make([]PanNode, initCap),
		nodeNum:   0,
		rootIndex: 0,
	}
}

func (tree *CTree) Add(data int) {

}

func tesAvl1() {
	root := &TreeNode{
		data:  3,
		left:  nil,
		right: nil,
	}
	root = root.avlTreeAdd(2)
	root = root.avlTreeAdd(1)
	preOrderTraversal(root)
	fmt.Println()
	root = root.avlTreeAdd(4)
	preOrderTraversal(root)
	fmt.Println()
	root = root.avlTreeAdd(5)
	preOrderTraversal(root)
	fmt.Println()
	root = root.avlTreeAdd(6)
	preOrderTraversal(root)
	fmt.Println("\n77777777777777777")
	root = root.avlTreeAdd(7)
	preOrderTraversal(root)
	fmt.Println("\n=====================")
	root = root.avlTreeAdd(10)
	preOrderTraversal(root)
	fmt.Println("\n-------------------------------")
	root = root.avlTreeAdd(9)
	preOrderTraversal(root)
	fmt.Println("\n--------------8888888888888888-----------------")
	root = root.avlTreeAdd(8)
	root.calTreeMinimumBalanceFactor()
	preOrderTraversal(root)
}

func tesAvl2() {
	root := &TreeNode{
		data:  7,
		left:  nil,
		right: nil,
	}
	root = root.avlTreeAdd(8)
	root = root.avlTreeAdd(3)
	preOrderTraversal(root)
	fmt.Println()
	root = root.avlTreeAdd(9)
	preOrderTraversal(root)
	fmt.Println()
	root = root.avlTreeAdd(5)
	preOrderTraversal(root)
	root = root.avlTreeAdd(4)
	preOrderTraversal(root)
}

/**
构建一个大顶堆(所有根节点的值大于左右孩子)
*/
func buildBigHeap(a []int) *TreeNode {

	if len(a) < 0 {
		return nil
	}
	tree := &TreeNode{
		data:          a[0],
		left:          nil,
		right:         nil,
		parent:        nil,
		balanceFactor: 0,
		nodeSize:      1,
	}

	tree = tree.bigHeapAddNode(tree, a[1])
	tree = tree.bigHeapAddNode(tree, a[2])
	tree = tree.bigHeapAddNode(tree, a[3])
	tree = tree.bigHeapAddNode(tree, a[4])
	tree = tree.bigHeapAddNode(tree, a[5])
	tree = tree.bigHeapAddNode(tree, 20)
	tree = tree.bigHeapAddNode(tree, 17)
	tree = tree.bigHeapAddNode(tree, 2)
	tree = tree.bigHeapAddNode(tree, 0)
	tree = tree.bigHeapAddNode(tree, 8)
	tree = tree.bigHeapAddNode(tree, 13)
	return tree
}

func (*TreeNode) bigHeapAddNode(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return &TreeNode{
			data:          data,
			left:          nil,
			right:         nil,
			parent:        nil,
			balanceFactor: 0,
			nodeSize:      1,
		}
	}
	tree, node := root.isFullBinaryTree(root)
	if !tree {
		if node.left == nil {
			node.left = node.left.bigHeapAddNode(node.left, data)
			node.left.parent = node
			left := node.left
			parent := left.parent
			if left.data < parent.data {
				parent.nodeSize++
			} else {
				for parent != nil && left.data > parent.data {
					tempData := parent.data
					parent.data = left.data
					left.data = tempData
					left = parent
					parent.nodeSize++
					parent = parent.parent
				}
			}
		} else {
			if node.right == nil {
				node.right = node.right.bigHeapAddNode(node.right, data)
				node.right.parent = node
				right := node.right
				parent := right.parent
				if right.data < parent.data {
					parent.nodeSize++
				} else {
					for parent != nil && right.data > parent.data {
						tempData := parent.data
						parent.data = right.data
						right.data = tempData
						right = parent
						parent.nodeSize++
						parent = parent.parent
					}
				}
			}
		}
	} else {
		root.left = root.left.bigHeapAddNode(root.left, data)
		root.nodeSize++
	}
	return root
}

func testBuildBigHeap() {
	a := []int{4, 5, 1, 3, 9, 11}
	heap := buildBigHeap(a)
	preOrderTraversal(heap)
	// fmt.Println()
	// middleOrderTraversal(heap)
	// tree := &TreeNode{
	// 	data:          5,
	// 	left:          nil,
	// 	right:         nil,
	// 	parent:        nil,
	// 	balanceFactor: 0,
	// 	nodeSize:      1,
	// }
	// tree.add(4)
	// tree.right = &TreeNode{
	// 	data:          1,
	// 	left:          nil,
	// 	right:         nil,
	// 	parent:        nil,
	// 	balanceFactor: 0,
	// 	nodeSize:      1,
	// }
	// tree.nodeSize++
	// tree.add(11)
	// tree.add(20)
	// tree.add(35)
	// tree.add(60)
	// tree.add(10)
	// tree.add(12)
	// tree.add(18)
	// tree.add(21)
	// // tree.add(25)
	// // tree.add(40)
	// binaryTree, node := tree.isFullBinaryTree(tree)
	// fmt.Printf("tree.isFullBinaryTree=%v node=%v\n", binaryTree, node)
	// if !binaryTree {
	// 	fmt.Printf("node.left=%v", node.left)
	// }
}
func main() {

	testBuildBigHeap()
	// tesAvl2()
	// testHuffmanTree()
	// root := &TreeNode{
	// 	data:  8,
	// 	left:  nil,
	// 	right: nil,
	// }
	// root.add(5)
	// root.add(15)
	// root.add(3)
	// root.add(7)
	// root.add(1)
	// root.add(11)
	// root.add(23)
	// root.add(9)
	// root.add(20)
	// root.add(21)
	//
	// fmt.Printf("postOrderTraversal\n")
	// postOrderTraversal(root)
	// fmt.Printf("\npreOrderTraversal\n")
	// preOrderTraversal(root)
	// fmt.Printf("\nmiddleOrderTraversal\n")
	// middleOrderTraversal(root)
	// fmt.Printf("\ntierOrderTraversal\n")
	// tierOrderTraversal(root)
	// node, contain := root.contains(5)
	// fmt.Printf("%d contains of = %v\n", 5, contain)
	// if contain {
	// 	fmt.Printf("%d.parent=%d\n", node.data, node.parent.data)
	// }
	// node, contain = root.contains(3)
	// fmt.Printf("%d contains of = %v\n", 3, contain)
	// if contain {
	// 	fmt.Printf("%d.parent=%d\n", node.data, node.parent.data)
	// }
	// node, contain = root.contains(7)
	// fmt.Printf("%d contains of = %v\n", 7, contain)
	// if contain {
	// 	fmt.Printf("%d.parent=%d\n", node.data, node.parent.data)
	// }
	// data := root.findMin()
	// fmt.Printf("min =%v \n", data.data)
	// fmt.Printf("%d.parent=%d\n", data.data, data.parent.data)

	// data = root.findMax()
	// fmt.Printf("Max =%v \n", data.data)
	// fmt.Printf("%d.parent=%d\n", data.data, data.parent.data)

	// println("remove====================")
	// targetNode := root.remove(root, 15)
	// middleOrderTraversal(targetNode)
	//
	// println()
	// node, contain := root.contains(21)
	// if contain && node.parent != nil {
	// 	fmt.Printf("%d.parent=%d", node.data, node.parent.data)
	// }
	//
	// println()
	// height := root.height()
	// fmt.Printf("height %d \n", height)
}
