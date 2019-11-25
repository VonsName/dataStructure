package main

import (
	"fmt"
	"regexp"
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

func (tree *TreeNode) add(data int) *TreeNode {
	if tree == nil {
		return &TreeNode{
			data:  data,
			left:  nil,
			right: nil,
		}
	}
	if data < tree.data {
		tree.left = tree.left.add(data)
		tree.left.parent = tree
	} else {
		tree.right = tree.right.add(data)
		tree.right.parent = tree
	}
	return tree
}

func (tree *TreeNode) contains(data int) (node *TreeNode, contain bool) {

	if tree == nil {
		return nil, false
	}
	if data == tree.data {
		return tree, true
	} else if data > tree.data {
		return tree.right.contains(data)
	} else if data < tree.data {
		return tree.left.contains(data)
	}
	return nil, false
}

/**
删除节点
*/
func (tree *TreeNode) remove(target int) (targetNode *TreeNode, notExists bool) {

	if tree == nil { // 空树
		return nil, true
	}
	if target == tree.data {
		if tree.left != nil && tree.right != nil {
			minNode := tree.right.findMin()
			minNode.parent.left = minNode.right
		}
	}

	return nil, true
}

func (tree *TreeNode) findMin() (minNode *TreeNode) {
	if tree == nil {
		return nil
	}
	if tree.left != nil {
		return tree.left.findMin()
	}
	return tree
}

func (tree *TreeNode) findMax() (maxNode *TreeNode) {
	if tree == nil {
		return nil
	}
	if tree.right != nil {
		return tree.right.findMax()
	}
	return tree
}

func main() {
	root := &TreeNode{
		data:  8,
		left:  nil,
		right: nil,
	}
	root.add(5)
	root.add(15)
	root.add(3)
	root.add(7)
	root.add(1)
	root.add(11)
	root.add(23)
	root.add(9)
	root.add(20)
	root.add(55)

	postOrderTraversal(root)
	println()
	preOrderTraversal(root)
	println()
	middleOrderTraversal(root)
	println()
	node, contain := root.contains(5)
	fmt.Printf("%d contains of = %v\n", 5, contain)
	if contain {
		fmt.Printf("%d.parent=%d\n", node.data, node.parent.data)
	}
	node, contain = root.contains(3)
	fmt.Printf("%d contains of = %v\n", 3, contain)
	if contain {
		fmt.Printf("%d.parent=%d\n", node.data, node.parent.data)
	}
	node, contain = root.contains(7)
	fmt.Printf("%d contains of = %v\n", 7, contain)
	if contain {
		fmt.Printf("%d.parent=%d\n", node.data, node.parent.data)
	}
	data := root.findMin()
	fmt.Printf("min =%v \n", data.data)
	fmt.Printf("%d.parent=%d\n", data.data, data.parent.data)

	data = root.findMax()
	fmt.Printf("Max =%v \n", data.data)
	fmt.Printf("%d.parent=%d\n", data.data, data.parent.data)
}
