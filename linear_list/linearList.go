package linear_list

import (
	"dataStructure/queue"
	"fmt"
	"strconv"
)

type MList interface {
	Add(Data int) int
	Remove(Data int)
	Clear()
	Show()
	RemoveByIndex(index int) int
	Exists(Data int) bool
	Size() int
	Get(index int) int
}

type SqList struct {
	Data []int
	size int
}

func newSqList(initialize int) *SqList {

	return &SqList{
		Data: make([]int, initialize),
		size: 0,
	}
}

func (list *SqList) Add(Data int) int {
	if list.size == len(list.Data) {
		panic("full index")
	}
	list.Data[list.size] = Data
	list.size++
	return Data
}

func (list *SqList) Size() int {
	return list.size
}

func (list *SqList) Get(index int) int {
	if index < 0 || index >= len(list.Data) {
		panic(fmt.Sprintf("outbound index of %d", index))
	}
	return list.Data[index]
}

func (list *SqList) Clear() {
	list.Data = make([]int, len(list.Data))
}

func (list *SqList) Remove(Data int) {
	if len(list.Data) == 0 {
		panic("list must not be nil")
	}
	if !list.Exists(Data) {
		panic(fmt.Sprintf("%d is not exists", Data))
	}
	index := 0
	for k, v := range list.Data {
		if v == Data {
			index = k
		}
	}
	list.Data[index] = 0
	for i := index; i < list.Size()-1; i++ {
		list.Data[i] = list.Data[i+1]
	}
	list.size--
}

func (list *SqList) Exists(Data int) (exist bool) {
	exist = false
	for i := 0; i < list.size; i++ {
		if list.Data[i] == Data {
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
	Data := list.Data[index]
	for i := index; i < list.Size(); i++ {
		list.Data[i] = list.Data[i+1]
	}
	list.size--
	return Data
}
func (list *SqList) Show() {

	for i := 0; i < list.size; i++ {
		fmt.Printf("k=%d,v=%d\n", i, list.Data[i])
	}
	fmt.Println()
}

type Node struct {
	Data     interface{}
	Next     *Node
	Previous *Node
}
type LinkedList struct {
	header *Node
	tail   *Node
	size   int
}

func newLinkedList() *LinkedList {
	header := &Node{
		Data: 0,
		Next: nil,
	}
	return &LinkedList{
		header: header,
		tail:   header,
	}
}

func (list *LinkedList) Add(Data int) int {
	list.size++
	if list.header.Next == nil {
		list.header.Data = Data
		list.header.Next = list.tail
		return Data
	} else {
		newNode := &Node{
			Data: Data,
			Next: nil,
		}
		if list.header == list.tail {
			list.tail = newNode
			list.header.Next = list.tail
		} else {
			list.tail.Next = newNode
			list.tail = newNode
		}
		return Data
	}
}

func (list *LinkedList) Show() {
	node := list.header
	for node != nil {
		fmt.Printf("%d ", node.Data)
		node = node.Next
	}
	fmt.Println()
}

func (list *LinkedList) Clear() {
	list.header.Data = 0
	list.header.Next = nil
	list.tail = list.header
	list.size = 0
}

func (list *LinkedList) Remove(Data int) {
	if !list.Exists(Data) {
		panic("元素不存在")
	}
	node := list.header
	list.size--
	if node.Data == Data {
		list.header = list.header.Next
		return
	}
	for node != nil && node.Next != nil {
		if node.Next.Data == Data {
			if node.Next == list.tail {
				list.tail = node
			}
			node.Next = node.Next.Next
		}
		node = node.Next
	}
}

func (list *LinkedList) Exists(Data int) (exist bool) {
	node := list.header
	for node != nil {
		if node.Data == Data {
			return true
		}
		node = node.Next
	}
	return false
}

func (list *LinkedList) Get(index int) int {
	if index < 0 || index > list.size {
		panic(fmt.Sprintf("index outbound of %d\n", index))
	}
	node := list.header
	for i := 0; i < index; i++ {
		node = node.Next
	}
	if node != nil {
		return node.Data.(int)
	}
	panic("outbound of index")
}

func (list *LinkedList) RemoveByIndex(index int) (Data int) {
	if index < 0 || index > list.size {
		panic(fmt.Sprintf("index outbound of %d\n", index))
	}
	node := list.header
	for i := 0; i < index-1; i++ {
		node = node.Next
	}
	if node == list.header {
		Data = list.header.Data.(int)
		list.header = list.header.Next
		list.size--
		return
	} else {
		Data = node.Next.Data.(int)
		if node.Next == list.tail {
			list.tail = node
		}
		node.Next = node.Next.Next
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
	if node.Next == nil {
		list.header = node
		list.tail = node
		return
	}
	reverses(list, node.Next)
	list.tail.Next = node
	list.tail = node
	list.tail.Next = nil
}

func (list *LinkedList) AddAll(targetList *LinkedList) {
	if targetList == nil {
		panic("targetList must not be nil")
	}
	list.tail.Next = targetList.header
	list.tail = targetList.tail
}

type CircleBothWayLinkedList struct {
	header *Node
	tail   *Node
	size   int
}

func newCircleLinkedList() *CircleBothWayLinkedList {

	header := &Node{
		Data:     0,
		Next:     nil,
		Previous: nil,
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
	list.tail.Previous = list.header
	list.header.Next = nil
	list.header.Previous = list.tail
	list.size = 0
	list.header.Data = -99999
}

func (list *CircleBothWayLinkedList) AddAll(targetList *CircleBothWayLinkedList) {
	if targetList == nil {
		panic("targetList must not be nil")
	}
	targetList.header.Previous = list.tail
	targetList.tail.Next = list.header
	list.tail.Next = targetList.header
	list.tail = targetList.tail
	list.header.Previous = targetList.tail
}

func (list *CircleBothWayLinkedList) Add(Data int) {
	if list.header.Next == nil {
		list.header.Data = Data
		list.header.Next = list.tail
	} else {
		node := &Node{
			Data:     Data,
			Next:     list.header,
			Previous: list.tail,
		}
		if list.header == list.tail {
			list.header.Next = node
		} else {
			list.tail.Next = node
		}
		list.tail = node
		list.header.Previous = list.tail
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
	for node.Next != list.header {
		fmt.Printf("%d ", node.Data)
		node = node.Next
	}
	fmt.Printf("%d \n", node.Data)
}

func (list *CircleBothWayLinkedList) Exists(Data int) (exist bool) {
	node := list.header
	for node.Next != list.header {
		if node.Data == Data {
			return true
		}
		node = node.Next
	}
	// 最后一个节点
	if node.Data == Data {
		return true
	}
	return false
}
func (list *CircleBothWayLinkedList) Remove(Data int) {
	if !list.Exists(Data) {
		panic(fmt.Sprintf("Data %d not exists\n", Data))
	}
	node := list.header
	for node.Next != list.header {
		if node.Data == Data {
			deletes(node, list)
			break
		}
		node = node.Next
	}
	// 最后一个尾结点的比较
	if node.Data == Data {
		deletes(node, list)
	}
}

func deletes(node *Node, list *CircleBothWayLinkedList) {
	if node == list.header {
		list.header = list.header.Next
		list.tail.Next = list.header
		list.header.Previous = list.tail
	} else if node == list.tail {
		list.tail = list.tail.Previous
		list.tail.Next = list.header
	} else {
		temp := node.Next
		temp.Previous = node.Previous
		node.Previous.Next = temp
	}
	list.size--
}

func (list *CircleBothWayLinkedList) RemoveByIndex(index int) int {
	if index < 0 || index >= list.size {
		panic(fmt.Sprintf("index %d out bound \n", index))
	}
	node := list.header
	for i := 0; i < index; i++ {
		node = node.Next
	}
	deletes(node, list)
	return node.Data.(int)
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
		node = node.Next
	}
	return node.Data.(int)
}

func (list *CircleBothWayLinkedList) Reverse() {
	r(list, list.header)
}

// 1 2 3 4 5
// 5 4 3 2 1
func r(list *CircleBothWayLinkedList, node *Node) {
	if node.Next == list.header {
		header := list.header
		list.header = node
		list.header.Previous = header
		list.tail = list.header
		return
	}
	r(list, node.Next)
	node.Previous = list.tail
	list.tail.Next = node
	list.tail = node
	list.tail.Next = list.header
}

type LinkedStack struct {
	top  *Node
	size int
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{top: &Node{
		Data:     0,
		Next:     nil,
		Previous: nil,
	}}
}

func (stack *LinkedStack) Push(Data interface{}) {

	if stack.size == 0 {
		if stack.top != nil {
			stack.top.Data = Data
		} else {
			stack.top = &Node{
				Data:     Data,
				Next:     nil,
				Previous: nil,
			}
		}
	} else {
		newNode := &Node{
			Data:     Data,
			Next:     nil,
			Previous: nil,
		}
		newNode.Next = stack.top
		stack.top = newNode
	}
	stack.size++
}

func (stack *LinkedStack) Pop() *Node {
	if stack.size == 0 {
		panic("stack is empty")
	}
	node := stack.top
	stack.top = stack.top.Next
	stack.size--
	return node
}

func (stack *LinkedStack) Show() {

	node := stack.top
	for node != nil {
		fmt.Printf("%v ", node.Data)
		node = node.Next
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
	stack := NewLinkedStack()
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
				for pop.Data.(string) != "(" {
					if err := circleQueue.EnQueue(pop.Data); err != nil {
						panic(err)
					}
					pop = stack.Pop()
				}
			} else {
				if stack.IsEmpty() {
					stack.Push(string(v))
				} else {
					Data := stack.top.Data
					// 区别对待 (
					if Data.(string) == "(" {
						stack.Push(string(v))
					} else {
						// 当前元素的优先级低于栈顶的元素的优先级,弹出并输出所有元素,然后将当前元素入栈
						if getPriority(string(v)) < getPriority(Data.(string)) {
							Data = stack.Pop().Data
							for getPriority(string(v)) < getPriority(Data.(string)) {
								if err := circleQueue.EnQueue(Data.(string)); err != nil {
									panic(err)
								}
								if stack.top.Data.(string) == "(" {
									break
								}
								Data = stack.Pop().Data
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
		if err := circleQueue.EnQueue(stack.Pop().Data.(string)); err != nil {
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
	stack := NewLinkedStack()
	for _, v := range expression {
		if isExpress(string(v)) {
			pop1, _ := strconv.Atoi(stack.Pop().Data.(string))
			pop2, _ := strconv.Atoi(stack.Pop().Data.(string))
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
	if i2, ok := stack.Pop().Data.(string); ok {
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
	stack := NewLinkedStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push("*")
	stack.Show()

	fmt.Printf("pop=%v\n", stack.Pop().Data)
	fmt.Printf("pop=%v\n", stack.Pop().Data)
}

func testCircleLinkedList() {

	list := newCircleLinkedList()
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Show()
	fmt.Printf("header:%d \n", list.header.Data)
	fmt.Printf("header Previous:%d \n", list.header.Previous.Data)
	fmt.Printf("header Next:%d \n", list.header.Next.Data)
	fmt.Printf("tail :%d \n", list.tail.Data)
	fmt.Printf("tail Next:%d \n", list.tail.Next.Data)
	fmt.Printf("tail Previous:%d \n", list.tail.Previous.Data)

	fmt.Println("\nrem!6666!ove!!!")
	list.Remove(6)
	list.Add(7)
	list.Show()
	fmt.Printf("header:%d \n", list.header.Data)
	fmt.Printf("header Previous:%d \n", list.header.Previous.Data)
	fmt.Printf("header Next:%d \n", list.header.Next.Data)
	fmt.Printf("\ntail :%d \n", list.tail.Data)
	fmt.Printf("tail Next:%d \n", list.tail.Next.Data)
	fmt.Printf("tail Previous:%d \n", list.tail.Previous.Data)

	fmt.Printf("\nrem22222ove \n")
	list.Remove(2)
	list.Add(10)
	list.Show()
	fmt.Printf("header:%d \n", list.header.Data)
	fmt.Printf("header Previous:%d \n", list.header.Previous.Data)
	fmt.Printf("header Next:%d \n", list.header.Next.Data)
	fmt.Printf("\ntail :%d \n", list.tail.Data)
	fmt.Printf("tail Next:%d \n", list.tail.Next.Data)
	fmt.Printf("tail Previous:%d \n", list.tail.Previous.Data)

	fmt.Printf("\nrem55555ove \n")
	list.Remove(5)
	list.Show()
	fmt.Printf("header:%d \n", list.header.Data)
	fmt.Printf("header Previous:%d \n", list.header.Previous.Data)
	fmt.Printf("header Next:%d \n", list.header.Next.Data)
	fmt.Printf("\ntail :%d \n", list.tail.Data)
	fmt.Printf("tail Next:%d \n", list.tail.Next.Data)
	fmt.Printf("tail Previous:%d \n", list.tail.Previous.Data)

	fmt.Printf("\nrem--ve by index\n")
	list.RemoveByIndex(1)
	list.Show()
	fmt.Printf("header:%d \n", list.header.Data)
	fmt.Printf("header Previous:%d \n", list.header.Previous.Data)
	fmt.Printf("header Next:%d \n", list.header.Next.Data)
	fmt.Printf("\ntail :%d \n", list.tail.Data)
	fmt.Printf("tail Next:%d \n", list.tail.Next.Data)
	fmt.Printf("tail Previous:%d \n", list.tail.Previous.Data)

	list.Add(33)
	list.Add(55)
	list.Show()

	list.Remove(10)
	list.Reverse()
	fmt.Printf("\nreverse1 after-----\n")
	list.Show()
	fmt.Printf("header:%d \n", list.header.Data)
	fmt.Printf("header Previous:%d \n", list.header.Previous.Data)
	fmt.Printf("header Next:%d \n", list.header.Next.Data)
	fmt.Printf("\ntail :%d \n", list.tail.Data)
	fmt.Printf("tail Next:%d \n", list.tail.Next.Data)
	fmt.Printf("tail Previous:%d \n", list.tail.Previous.Data)

	list.Reverse()
	fmt.Printf("\nreverse2 after-----\n")
	list.Show()
	fmt.Printf("header:%d \n", list.header.Data)
	fmt.Printf("header Previous:%d \n", list.header.Previous.Data)
	fmt.Printf("header Next:%d \n", list.header.Next.Data)
	fmt.Printf("\ntail :%d \n", list.tail.Data)
	fmt.Printf("tail Next:%d \n", list.tail.Next.Data)
	fmt.Printf("tail Previous:%d \n", list.tail.Previous.Data)

	fmt.Printf("\nadd all after-----\n")
	linkedList := newCircleLinkedList()
	linkedList.Add(99)
	linkedList.Add(100)
	linkedList.Add(88)
	linkedList.Add(66)
	linkedList.Add(77)
	list.AddAll(linkedList)
	list.Show()
	fmt.Printf("header:%d \n", list.header.Data)
	fmt.Printf("header Previous:%d \n", list.header.Previous.Data)
	fmt.Printf("header Next:%d \n", list.header.Next.Data)
	fmt.Printf("\ntail :%d \n", list.tail.Data)
	fmt.Printf("tail Next:%d \n", list.tail.Next.Data)
	fmt.Printf("tail Previous:%d \n", list.tail.Previous.Data)
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
	// Data := list.get(2)
	// fmt.Printf("get Data:%d\n", Data)
}

func testLinkedList() {
	linkedList := newLinkedList()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(5)
	linkedList.Show()
	fmt.Printf("tail=%d\n", linkedList.tail.Data)
	fmt.Printf("header=%d\n", linkedList.header.Data)

	linkedList.Remove(1)
	linkedList.Show()
	fmt.Printf("tail=%d\n", linkedList.tail.Data)
	fmt.Printf("header=%d\n", linkedList.header.Data)
	linkedList.Remove(5)
	linkedList.Show()
	fmt.Printf("tail=%d\n", linkedList.tail.Data)
	fmt.Printf("header=%d\n", linkedList.header.Data)
	fmt.Println()
	linkedList.Add(12)
	linkedList.Add(13)
	linkedList.Add(14)
	linkedList.Add(15)
	linkedList.Show()

	fmt.Printf("get %d Data=%d\n", 0, linkedList.Get(0))
	Data := linkedList.RemoveByIndex(0)
	fmt.Printf("remove Data=%d header=%d size=%d \n", Data, linkedList.header.Data, linkedList.size)
	linkedList.Show()

	Data = linkedList.RemoveByIndex(5)
	fmt.Printf("remove Data=%d tail=%d size=%d \n", Data, linkedList.tail.Data, linkedList.size)
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
	fmt.Printf("tail=%d\n", linkedList.tail.Data)
}
