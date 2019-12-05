package main

import (
	"dataStructure/queue"
	"fmt"
	"strconv"
)

type MList interface {
	Add(data int) int
	Remove(data int)
	Clear()
	Show()
	RemoveByIndex(index int) int
	Exists(data int) bool
	Size() int
	Get(index int) int
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

func (list *SqList) Get(index int) int {
	if index < 0 || index >= len(list.data) {
		panic(fmt.Sprintf("outbound index of %d", index))
	}
	return list.data[index]
}

func (list *SqList) Clear() {
	list.data = make([]int, len(list.data))
}

func (list *SqList) Remove(data int) {
	if len(list.data) == 0 {
		panic("list must not be nil")
	}
	if !list.Exists(data) {
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

func (list *SqList) Exists(data int) (exist bool) {
	exist = false
	for i := 0; i < list.size; i++ {
		if list.data[i] == data {
			exist = true
			break
		}
	}
	return
}
func (list *SqList) RemoveByIndex(index int) int {
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
func (list *SqList) Show() {

	for i := 0; i < list.size; i++ {
		fmt.Printf("k=%d,v=%d\n", i, list.data[i])
	}
	fmt.Println()
}

type Node struct {
	data     interface{}
	next     *Node
	previous *Node
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

func (list *LinkedList) Show() {
	node := list.header
	for node != nil {
		fmt.Printf("%d ", node.data)
		node = node.next
	}
	fmt.Println()
}

func (list *LinkedList) Clear() {
	list.header.data = 0
	list.header.next = nil
	list.tail = list.header
	list.size = 0
}

func (list *LinkedList) Remove(data int) {
	if !list.Exists(data) {
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

func (list *LinkedList) Exists(data int) (exist bool) {
	node := list.header
	for node != nil {
		if node.data == data {
			return true
		}
		node = node.next
	}
	return false
}

func (list *LinkedList) Get(index int) int {
	if index < 0 || index > list.size {
		panic(fmt.Sprintf("index outbound of %d\n", index))
	}
	node := list.header
	for i := 0; i < index; i++ {
		node = node.next
	}
	if node != nil {
		return node.data.(int)
	}
	panic("outbound of index")
}

func (list *LinkedList) RemoveByIndex(index int) (data int) {
	if index < 0 || index > list.size {
		panic(fmt.Sprintf("index outbound of %d\n", index))
	}
	node := list.header
	for i := 0; i < index-1; i++ {
		node = node.next
	}
	if node == list.header {
		data = list.header.data.(int)
		list.header = list.header.next
		list.size--
		return
	} else {
		data = node.next.data.(int)
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

func (list *LinkedList) Reverse() {
	reverses(list, list.header)
}
func reverses(list *LinkedList, node *Node) {
	if node.next == nil {
		list.header = node
		list.tail = node
		return
	}
	reverses(list, node.next)
	list.tail.next = node
	list.tail = node
	list.tail.next = nil
}

func (list *LinkedList) AddAll(targetList *LinkedList) {
	if targetList == nil {
		panic("targetList must not be nil")
	}
	list.tail.next = targetList.header
	list.tail = targetList.tail
}

type CircleBothWayLinkedList struct {
	header *Node
	tail   *Node
	size   int
}

func newCircleLinkedList() *CircleBothWayLinkedList {

	header := &Node{
		data:     0,
		next:     nil,
		previous: nil,
	}
	list := &CircleBothWayLinkedList{
		header: header,
		tail:   header,
		size:   0,
	}
	initCircleLinkedList(list)
	return list
}

func initCircleLinkedList(list *CircleBothWayLinkedList) {
	list.tail = list.header
	list.tail.previous = list.header
	list.header.next = nil
	list.header.previous = list.tail
	list.size = 0
	list.header.data = -99999
}

func (list *CircleBothWayLinkedList) AddAll(targetList *CircleBothWayLinkedList) {
	if targetList == nil {
		panic("targetList must not be nil")
	}
	targetList.header.previous = list.tail
	targetList.tail.next = list.header
	list.tail.next = targetList.header
	list.tail = targetList.tail
	list.header.previous = targetList.tail
}

func (list *CircleBothWayLinkedList) Add(data int) {
	if list.header.next == nil {
		list.header.data = data
		list.header.next = list.tail
	} else {
		node := &Node{
			data:     data,
			next:     list.header,
			previous: list.tail,
		}
		if list.header == list.tail {
			list.header.next = node
		} else {
			list.tail.next = node
		}
		list.tail = node
		list.header.previous = list.tail
	}
	list.size++
}

func (list *CircleBothWayLinkedList) Size() int {
	return list.size
}

func (list *CircleBothWayLinkedList) Show() {
	if list.size == 0 {
		return
	}
	node := list.header
	for node.next != list.header {
		fmt.Printf("%d ", node.data)
		node = node.next
	}
	fmt.Printf("%d \n", node.data)
}

func (list *CircleBothWayLinkedList) Exists(data int) (exist bool) {
	node := list.header
	for node.next != list.header {
		if node.data == data {
			return true
		}
		node = node.next
	}
	// 最后一个节点
	if node.data == data {
		return true
	}
	return false
}
func (list *CircleBothWayLinkedList) Remove(data int) {
	if !list.Exists(data) {
		panic(fmt.Sprintf("data %d not exists\n", data))
	}
	node := list.header
	for node.next != list.header {
		if node.data == data {
			deletes(node, list)
			break
		}
		node = node.next
	}
	// 最后一个尾结点的比较
	if node.data == data {
		deletes(node, list)
	}
}

func deletes(node *Node, list *CircleBothWayLinkedList) {
	if node == list.header {
		list.header = list.header.next
		list.tail.next = list.header
		list.header.previous = list.tail
	} else if node == list.tail {
		list.tail = list.tail.previous
		list.tail.next = list.header
	} else {
		temp := node.next
		temp.previous = node.previous
		node.previous.next = temp
	}
	list.size--
}

func (list *CircleBothWayLinkedList) RemoveByIndex(index int) int {
	if index < 0 || index >= list.size {
		panic(fmt.Sprintf("index %d out bound \n", index))
	}
	node := list.header
	for i := 0; i < index; i++ {
		node = node.next
	}
	deletes(node, list)
	return node.data.(int)
}
func (list *CircleBothWayLinkedList) Clear() {
	initCircleLinkedList(list)
}
func (list *CircleBothWayLinkedList) Get(index int) int {
	if index < 0 || index >= list.size {
		panic(fmt.Sprintf("index %d out bound \n", index))
	}
	node := list.header
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.data.(int)
}

func (list *CircleBothWayLinkedList) Reverse() {
	r(list, list.header)
}

// 1 2 3 4 5
// 5 4 3 2 1
func r(list *CircleBothWayLinkedList, node *Node) {
	if node.next == list.header {
		header := list.header
		list.header = node
		list.header.previous = header
		list.tail = list.header
		return
	}
	r(list, node.next)
	node.previous = list.tail
	list.tail.next = node
	list.tail = node
	list.tail.next = list.header
}

type LinkedStack struct {
	top  *Node
	size int
}

func newLinkedStack() *LinkedStack {
	return &LinkedStack{top: &Node{
		data:     0,
		next:     nil,
		previous: nil,
	}}
}

func (stack *LinkedStack) Push(data interface{}) {

	if stack.size == 0 {
		if stack.top != nil {
			stack.top.data = data
		} else {
			stack.top = &Node{
				data:     data,
				next:     nil,
				previous: nil,
			}
		}
	} else {
		newNode := &Node{
			data:     data,
			next:     nil,
			previous: nil,
		}
		newNode.next = stack.top
		stack.top = newNode
	}
	stack.size++
}

func (stack *LinkedStack) Pop() *Node {
	if stack.size == 0 {
		panic("stack is empty")
	}
	node := stack.top
	stack.top = stack.top.next
	stack.size--
	return node
}

func (stack *LinkedStack) Show() {

	node := stack.top
	for node != nil {
		fmt.Printf("%v ", node.data)
		node = node.next
	}
	fmt.Println()
}

func (stack *LinkedStack) IsEmpty() (empty bool) {
	return stack.size == 0
}

// 9 + ( 3 - 1 ) * 3 + 10 / 2
// "((9 + ( 3 - 1 ))* 3 + 6) / 2"
// 中缀表达式转为后缀表达式
func generateSuffixExpression(str string) (expression string) {
	stack := newLinkedStack()
	circleQueue := queue.NewCircleQueue(len(str))
	for _, v := range str {
		if string(v) == " " {
			continue
		}
		if isExpress(string(v)) {
			if string(v) == "(" {
				stack.Push(string(v))
			} else if string(v) == ")" {
				pop := stack.Pop()
				for pop.data.(string) != "(" {
					if err := circleQueue.EnQueue(pop.data); err != nil {
						panic(err)
					}
					pop = stack.Pop()
				}
			} else {
				if stack.IsEmpty() {
					stack.Push(string(v))
				} else {
					data := stack.top.data
					// 区别对待 (
					if data.(string) == "(" {
						stack.Push(string(v))
					} else {
						// 当前元素的优先级低于栈顶的元素的优先级,弹出并输出所有元素,然后将当前元素入栈
						if getPriority(string(v)) < getPriority(data.(string)) {
							data = stack.Pop().data
							for getPriority(string(v)) < getPriority(data.(string)) {
								if err := circleQueue.EnQueue(data.(string)); err != nil {
									panic(err)
								}
								if stack.top.data.(string) == "(" {
									break
								}
								data = stack.Pop().data
							}
							stack.Push(string(v))
						} else { // 高于,直接入栈
							stack.Push(string(v))
						}
					}
				}
			}
		} else {
			if err := circleQueue.EnQueue(string(v)); err != nil {
				panic(err)
			}
		}
	}
	for !stack.IsEmpty() {
		if err := circleQueue.EnQueue(stack.Pop().data.(string)); err != nil {
			panic(err)
		}
	}
	circleQueue.Show()
	for circleQueue.Size() != 0 {
		deQueue, _ := circleQueue.DeQueue()
		expression += fmt.Sprintf("%s", deQueue)
	}
	return
}

// 计算后缀表达式
func calculateSuffixExpression(expression string) int {
	stack := newLinkedStack()
	for _, v := range expression {
		if isExpress(string(v)) {
			pop1, _ := strconv.Atoi(stack.Pop().data.(string))
			pop2, _ := strconv.Atoi(stack.Pop().data.(string))
			s := string(v)
			switch s {
			case "*":
				stack.Push(strconv.Itoa(pop2 * pop1))
			case "/":
				stack.Push(strconv.Itoa(pop2 / pop1))
			case "+":
				stack.Push(strconv.Itoa(pop2 + pop1))
			case "-":
				stack.Push(strconv.Itoa(pop2 - pop1))
			default:
				panic("invalid expression error")
			}
		} else {
			stack.Push(string(v))
		}
	}
	if i2, ok := stack.Pop().data.(string); ok {
		i, _ := strconv.Atoi(i2)
		return i
	} else {
		panic("Atoi error")
	}
}

func isExpress(c string) (ok bool) {

	switch c {
	case "+", "-", "*", "/", "(", ")":
		return true
	}
	return false
}

func getPriority(b string) int {
	switch b {
	case "+", "-":
		return 0
	case "*", "/":
		return 1
	case "(":
		return 2
	default:
		panic("express error")
	}
}

func main() {
	// testSqList()
	// testLinkedList()
	// testCircleLinkedList()
	// testLinkedStack()
	testSuffix()

}

func testSuffix() {
	str := "((9 + ( 3 - 1 ))* 3 + 6) / 2"
	expression := generateSuffixExpression(str)
	fmt.Printf("expression=%s\n", expression)

	suffixExpression := calculateSuffixExpression(expression)
	fmt.Printf("result=%d \n", suffixExpression)
}

func testLinkedStack() {
	stack := newLinkedStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push("*")
	stack.Show()

	fmt.Printf("pop=%v\n", stack.Pop().data)
	fmt.Printf("pop=%v\n", stack.Pop().data)
}

func testCircleLinkedList() {

	list := newCircleLinkedList()
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Show()
	fmt.Printf("header:%d \n", list.header.data)
	fmt.Printf("header previous:%d \n", list.header.previous.data)
	fmt.Printf("header next:%d \n", list.header.next.data)
	fmt.Printf("tail :%d \n", list.tail.data)
	fmt.Printf("tail next:%d \n", list.tail.next.data)
	fmt.Printf("tail previous:%d \n", list.tail.previous.data)

	fmt.Println("\nrem!6666!ove!!!")
	list.Remove(6)
	list.Add(7)
	list.Show()
	fmt.Printf("header:%d \n", list.header.data)
	fmt.Printf("header previous:%d \n", list.header.previous.data)
	fmt.Printf("header next:%d \n", list.header.next.data)
	fmt.Printf("\ntail :%d \n", list.tail.data)
	fmt.Printf("tail next:%d \n", list.tail.next.data)
	fmt.Printf("tail previous:%d \n", list.tail.previous.data)

	fmt.Printf("\nrem22222ove \n")
	list.Remove(2)
	list.Add(10)
	list.Show()
	fmt.Printf("header:%d \n", list.header.data)
	fmt.Printf("header previous:%d \n", list.header.previous.data)
	fmt.Printf("header next:%d \n", list.header.next.data)
	fmt.Printf("\ntail :%d \n", list.tail.data)
	fmt.Printf("tail next:%d \n", list.tail.next.data)
	fmt.Printf("tail previous:%d \n", list.tail.previous.data)

	fmt.Printf("\nrem55555ove \n")
	list.Remove(5)
	list.Show()
	fmt.Printf("header:%d \n", list.header.data)
	fmt.Printf("header previous:%d \n", list.header.previous.data)
	fmt.Printf("header next:%d \n", list.header.next.data)
	fmt.Printf("\ntail :%d \n", list.tail.data)
	fmt.Printf("tail next:%d \n", list.tail.next.data)
	fmt.Printf("tail previous:%d \n", list.tail.previous.data)

	fmt.Printf("\nrem--ve by index\n")
	list.RemoveByIndex(1)
	list.Show()
	fmt.Printf("header:%d \n", list.header.data)
	fmt.Printf("header previous:%d \n", list.header.previous.data)
	fmt.Printf("header next:%d \n", list.header.next.data)
	fmt.Printf("\ntail :%d \n", list.tail.data)
	fmt.Printf("tail next:%d \n", list.tail.next.data)
	fmt.Printf("tail previous:%d \n", list.tail.previous.data)

	list.Add(33)
	list.Add(55)
	list.Show()

	list.Remove(10)
	list.Reverse()
	fmt.Printf("\nreverse1 after-----\n")
	list.Show()
	fmt.Printf("header:%d \n", list.header.data)
	fmt.Printf("header previous:%d \n", list.header.previous.data)
	fmt.Printf("header next:%d \n", list.header.next.data)
	fmt.Printf("\ntail :%d \n", list.tail.data)
	fmt.Printf("tail next:%d \n", list.tail.next.data)
	fmt.Printf("tail previous:%d \n", list.tail.previous.data)

	list.Reverse()
	fmt.Printf("\nreverse2 after-----\n")
	list.Show()
	fmt.Printf("header:%d \n", list.header.data)
	fmt.Printf("header previous:%d \n", list.header.previous.data)
	fmt.Printf("header next:%d \n", list.header.next.data)
	fmt.Printf("\ntail :%d \n", list.tail.data)
	fmt.Printf("tail next:%d \n", list.tail.next.data)
	fmt.Printf("tail previous:%d \n", list.tail.previous.data)

	fmt.Printf("\nadd all after-----\n")
	linkedList := newCircleLinkedList()
	linkedList.Add(99)
	linkedList.Add(100)
	linkedList.Add(88)
	linkedList.Add(66)
	linkedList.Add(77)
	list.AddAll(linkedList)
	list.Show()
	fmt.Printf("header:%d \n", list.header.data)
	fmt.Printf("header previous:%d \n", list.header.previous.data)
	fmt.Printf("header next:%d \n", list.header.next.data)
	fmt.Printf("\ntail :%d \n", list.tail.data)
	fmt.Printf("tail next:%d \n", list.tail.next.data)
	fmt.Printf("tail previous:%d \n", list.tail.previous.data)
}
func testSqList() {
	list := newSqList(5)
	list.Add(13)
	list.Add(22)
	list.Add(31)
	list.Add(99)
	list.Add(100)
	list.Show()
	list.Remove(22)
	list.Show()

	fmt.Println("removeByIndex")
	list.RemoveByIndex(1)
	list.Show()
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
	linkedList.Show()
	fmt.Printf("tail=%d\n", linkedList.tail.data)
	fmt.Printf("header=%d\n", linkedList.header.data)

	linkedList.Remove(1)
	linkedList.Show()
	fmt.Printf("tail=%d\n", linkedList.tail.data)
	fmt.Printf("header=%d\n", linkedList.header.data)
	linkedList.Remove(5)
	linkedList.Show()
	fmt.Printf("tail=%d\n", linkedList.tail.data)
	fmt.Printf("header=%d\n", linkedList.header.data)
	fmt.Println()
	linkedList.Add(12)
	linkedList.Add(13)
	linkedList.Add(14)
	linkedList.Add(15)
	linkedList.Show()

	fmt.Printf("get %d data=%d\n", 0, linkedList.Get(0))
	data := linkedList.RemoveByIndex(0)
	fmt.Printf("remove data=%d header=%d size=%d \n", data, linkedList.header.data, linkedList.size)
	linkedList.Show()

	data = linkedList.RemoveByIndex(5)
	fmt.Printf("remove data=%d tail=%d size=%d \n", data, linkedList.tail.data, linkedList.size)
	linkedList.Show()
	linkedList.Add(999)
	linkedList.Show()

	fmt.Printf("reverse1\n")
	linkedList.Reverse()
	linkedList.Show()

	fmt.Printf("reverse2\n")
	linkedList.Reverse()
	linkedList.Show()

	fmt.Printf("\nafter add all-------------\n")
	list := newLinkedList()
	list.Add(100)
	list.Add(99)
	list.Add(88)
	list.Add(77)
	linkedList.AddAll(list)
	linkedList.Show()
	fmt.Printf("tail=%d\n", linkedList.tail.data)
}
