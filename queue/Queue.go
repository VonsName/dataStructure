package queue

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	size  int
	front int
	back  int
	data  []interface{}
}

func NewCircleQueue(initCap int) *CircleQueue {

	queue := &CircleQueue{
		size:  0,
		front: 0,
		back:  0,
		data:  make([]interface{}, initCap),
	}
	return queue
}

func (cq *CircleQueue) EnQueue(d interface{}) (err error) {
	if cq.size == len(cq.data) {
		return fmt.Errorf("%s", errors.New("当前队列已满"))
	}
	cq.data[cq.back] = d
	cq.back = (cq.back + 1) % len(cq.data)
	cq.size++
	return nil
}

func (cq *CircleQueue) DeQueue() (interface{}, error) {
	if cq.size == 0 {
		return 0, fmt.Errorf("%s", errors.New("当前队列已空"))
	}
	d := cq.data[cq.front]
	cq.front = (cq.front + 1) % len(cq.data)
	cq.size--
	return d, nil
}

func (cq *CircleQueue) Size() int {
	return cq.size
}

func (cq *CircleQueue) Show() {
	index := cq.front
	for i := 0; i < cq.size; i++ {
		fmt.Printf("%v\n", cq.data[index])
		index = (index + 1) % len(cq.data)
	}
	fmt.Println()
}

func main() {

	queue := NewCircleQueue(6)
	err := queue.EnQueue(1)
	err = queue.EnQueue(2)
	err = queue.EnQueue(3)
	err = queue.EnQueue(90)
	// err = queue.EnQueue(100)
	err = queue.EnQueue("+")
	if err != nil {
		panic(err)
	}
	queue.Show()

	err = queue.EnQueue(6)
	// err = queue.EnQueue(7)
	//
	// queue.Show()
	// err = queue.EnQueue(77)
	// err = queue.EnQueue(9)
	// err = queue.EnQueue(10)
	// err = queue.EnQueue(11)
	// err = queue.EnQueue(12)
	// queue.show()
	_, _ = queue.DeQueue()
	_, _ = queue.DeQueue()
	_ = queue.EnQueue(90)
	_ = queue.EnQueue(56)
	fmt.Printf("%s\n", "==================")
	queue.Show()

	_, _ = queue.DeQueue()
	_, _ = queue.DeQueue()
	fmt.Printf("%s\n", "==================")
	queue.Show()
	_, _ = queue.DeQueue()
	_, _ = queue.DeQueue()
	_, _ = queue.DeQueue()
	_, _ = queue.DeQueue()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", "==================")
	queue.Show()
}
