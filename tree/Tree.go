package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"sort"
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

func (s *Stack) push(data interface{}) {
	s.top++
	s.data = append(s.data, data)
}

func (s *Stack) pop() (data interface{}) {
	if s.top == -1 {
		return
	}
	data = s.data[s.top]
	s.top--
	return
}

func (s *Stack) show() {
	i := s.top
	for i != -1 {
		fmt.Printf("%v ", s.data[i])
		i--
	}
	fmt.Println()
}

type TreeNode struct {
	data   int
	left   *TreeNode // 左子树
	right  *TreeNode // 右子树
	parent *TreeNode // 父节点
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
	fmt.Printf("node=%v ", tree.data)
}

/**
先序遍历
*/
func preOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Printf("node=%v ", tree.data)
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
	if tree.left != nil {
		middleOrderTraversal(tree.left)
	}
	fmt.Printf("node=%v ", tree.data)
	if tree.right != nil {
		middleOrderTraversal(tree.right)
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

func NewHuffmanTree() (huffmanTree HuffmanTree) {

	huffmanTree = HuffmanTree{}
	for i := 0; i < len(huffmanTree); i++ {
		rand.Seed(time.Now().UnixNano())

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
	temp := tree[i]
	tree[i].weight = tree[j].weight
	tree[j].weight = temp.weight
}
func initWeight(tree *HuffmanTree) {
	if tree == nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		tree[i].weight = rand.Intn(100)
	}
	sort.Sort(tree)
}

func select2MinWeight(k int) {

	for i := 0; i < k; i++ {

	}
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
	if root == nil {
		return &TreeNode{
			data:  data,
			left:  nil,
			right: nil,
		}
	}
	if data < root.data {
		root.left = root.left.add(data)
		root.left.parent = root
	} else {
		root.right = root.right.add(data)
		root.right.parent = root
	}
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
		return -1
	}
	return int(1 + math.Max(float64(root.left.height()), float64(root.right.height())))
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

func main() {

	testHuffmanTree()
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
	// postOrderTraversal(root)
	// println()
	// preOrderTraversal(root)
	// println()
	// middleOrderTraversal(root)
	// println()
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
