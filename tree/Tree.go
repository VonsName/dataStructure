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
	data  int
	left  *TreeNode // 左子树
	right *TreeNode // 右子树
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
	} else {
		root.right = root.right.add(data)
	}
	return root
}

func main() {
	root := &TreeNode{
		data:  3,
		left:  nil,
		right: nil,
	}
	root.add(5)
	root.add(2)
	root.add(7)

	postOrderTraversal(root)
	println()
	preOrderTraversal(root)
}
