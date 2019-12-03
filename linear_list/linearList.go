package main

import (
	"fmt"
)

type List interface {
	Add(data int) int
	remove(data int)
	clear()
	show()
	removeByIndex(index int) int
	exists(data int) bool
	Size() int
	get(index int) int
}

type SqList struct {
	data []int
	size int
}

func newSqList(initialize int) *SqList {

	return &SqList{
		data: make([]int, initialize),
		size: 0,
	}
}

func (list *SqList) Add(data int) int {
	if list.size == len(list.data) {
		panic("full index")
	}
	list.data[list.size] = data
	list.size++
	return data
}

func (list *SqList) Size() int {
	return list.size
}

func (list *SqList) get(index int) int {
	if index < 0 || index >= len(list.data) {
		panic(fmt.Sprintf("outbound index of %d", index))
	}
	return list.data[index]
}

func (list *SqList) clear() {
	list.data = make([]int, len(list.data))
}

func (list *SqList) remove(data int) {
	if len(list.data) == 0 {
		panic("list must not be nil")
	}
	if !list.exists(data) {
		panic(fmt.Sprintf("%d is not exists", data))
	}
	index := 0
	for k, v := range list.data {
		if v == data {
			index = k
		}
	}
	list.data[index] = 0
	for i := index; i < list.Size()-1; i++ {
		list.data[i] = list.data[i+1]
	}
	list.size--
}

func (list *SqList) exists(data int) (exist bool) {
	exist = false
	for i := 0; i < list.size; i++ {
		if list.data[i] == data {
			exist = true
			break
		}
	}
	return
}
func (list *SqList) removeByIndex(index int) int {
	if index < 0 || index > list.Size() {
		panic(fmt.Sprintf("index outbound of %d", index))
	}
	data := list.data[index]
	for i := index; i < list.Size(); i++ {
		list.data[i] = list.data[i+1]
	}
	list.size--
	return data
}
func (list *SqList) show() {

	for i := 0; i < list.size; i++ {
		fmt.Printf("k=%d,v=%d\n", i, list.data[i])
	}
	fmt.Println()
}

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	header *Node
	tail   *Node
	size   int
}

func newLinkedList() *LinkedList {
	header := &Node{
		data: 0,
		next: nil,
	}
	return &LinkedList{
		header: header,
		tail:   header,
	}
}

func (list *LinkedList) Add(data int) int {
	list.size++
	if list.header.next == nil {
		list.header.data = data
		list.header.next = list.tail
		return data
	} else {
		newNode := &Node{
			data: data,
			next: nil,
		}
		if list.header == list.tail {
			list.tail = newNode
			list.header.next = list.tail
		} else {
			list.tail.next = newNode
			list.tail = newNode
		}
		return data
	}
}

func (list *LinkedList) show() {
	node := list.header
	for node != nil {
		fmt.Printf("%d ", node.data)
		node = node.next
	}
	fmt.Println()
}

func (list *LinkedList) clear() {
	list.header.data = 0
	list.header.next = nil
	list.tail = list.header
	list.size = 0
}

func (list *LinkedList) remove(data int) {
	if !list.exists(data) {
		panic("元素不存在")
	}
	node := list.header
	list.size--
	if node.data == data {
		list.header = list.header.next
		return
	}
	for node != nil && node.next != nil {
		if node.next.data == data {
			if node.next == list.tail {
				list.tail = node
			}
			node.next = node.next.next
		}
		node = node.next
	}
}

func (list *LinkedList) exists(data int) (exist bool) {
	node := list.header
	for node != nil {
		if node.data == data {
			return true
		}
		node = node.next
	}
	return false
}

func (list *LinkedList) get(index int) int {
	if index < 0 || index > list.size {
		panic(fmt.Sprintf("index outbound of %d\n", index))
	}
	node := list.header
	for i := 0; i < index; i++ {
		node = node.next
	}
	if node != nil {
		return node.data
	}
	panic("outbound of index")
}

func (list *LinkedList) removeByIndex(index int) (data int) {
	if index < 0 || index > list.size {
		panic(fmt.Sprintf("index outbound of %d\n", index))
	}
	node := list.header
	for i := 0; i < index-1; i++ {
		node = node.next
	}
	if node == list.header {
		data = list.header.data
		list.header = list.header.next
		list.size--
		return
	} else {
		data = node.next.data
		if node.next == list.tail {
			list.tail = node
		}
		node.next = node.next.next
		list.size--
		return
	}
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) reverse() {
	reverses(list, list.header, 0)
}
func reverses(list *LinkedList, node *Node, i int) {
	if node.next == nil {
		return
	}
	i++
	reverses(list, node.next, i)
	if i == list.size {
		list.header = node
		list.tail = node
	}
	if i == 1 {
		list.tail = node
	}
}

type CircleLinkedList struct {
}

func main() {
	// testSqList()
	testLinkedList()
}

func testSqList() {
	list := newSqList(5)
	list.Add(13)
	list.Add(22)
	list.Add(31)
	list.Add(99)
	list.Add(100)
	list.show()
	list.remove(22)
	list.show()

	fmt.Println("removeByIndex")
	list.removeByIndex(1)
	list.show()
	// list.Add(55)
	// list.Add(66)
	// list.show()
	//
	// data := list.get(2)
	// fmt.Printf("get data:%d\n", data)
}

func testLinkedList() {
	linkedList := newLinkedList()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(5)
	linkedList.show()
	fmt.Printf("tail=%d\n", linkedList.tail.data)
	fmt.Printf("header=%d\n", linkedList.header.data)

	linkedList.remove(1)
	linkedList.show()
	fmt.Printf("tail=%d\n", linkedList.tail.data)
	fmt.Printf("header=%d\n", linkedList.header.data)
	linkedList.remove(5)
	linkedList.show()
	fmt.Printf("tail=%d\n", linkedList.tail.data)
	fmt.Printf("header=%d\n", linkedList.header.data)
	fmt.Println()
	linkedList.Add(12)
	linkedList.Add(13)
	linkedList.Add(14)
	linkedList.Add(15)
	linkedList.show()

	fmt.Printf("get %d data=%d\n", 0, linkedList.get(0))
	data := linkedList.removeByIndex(0)
	fmt.Printf("remove data=%d header=%d size=%d \n", data, linkedList.header.data, linkedList.size)
	linkedList.show()

	data = linkedList.removeByIndex(5)
	fmt.Printf("remove data=%d tail=%d size=%d \n", data, linkedList.tail.data, linkedList.size)
	linkedList.show()
	linkedList.Add(999)
	linkedList.show()
}
