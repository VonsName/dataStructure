package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	size  int
	front int
	back  int
	data  []int
}

func NewCircleQueue(initCap int) *CircleQueue {

	queue := &CircleQueue{
		size:  0,
		front: 0,
		back:  0,
		data:  make([]int, initCap),
	}
	return queue
}

func (cq *CircleQueue) enQueue(d int) (err error) {
	if cq.size == len(cq.data) {
		return fmt.Errorf("%s", errors.New("当前队列已满"))
	}
	cq.data[cq.back] = d
	cq.back = (cq.back + 1) % len(cq.data)
	cq.size++
	return nil
}

func (cq *CircleQueue) deQueue() (int, error) {
	if cq.size == 0 {
		return 0, fmt.Errorf("%s", errors.New("当前队列已空"))
	}
	d := cq.data[cq.front]
	cq.front = (cq.front + 1) % len(cq.data)
	cq.size--
	return d, nil
}

func (cq *CircleQueue) show() {
	index := cq.front
	for i := 0; i < cq.size; i++ {
		fmt.Printf("%d\n", cq.data[index])
		index = (index + 1) % len(cq.data)
	}
}

func main() {

	queue := NewCircleQueue(6)
	err := queue.enQueue(1)
	err = queue.enQueue(2)
	err = queue.enQueue(3)
	err = queue.enQueue(90)
	err = queue.enQueue(100)
	if err != nil {
		panic(err)
	}
	queue.show()

	err = queue.enQueue(6)
	//err = queue.enQueue(7)
	//
	//queue.show()
	//err = queue.enQueue(77)
	//err = queue.enQueue(9)
	//err = queue.enQueue(10)
	//err = queue.enQueue(11)
	//err = queue.enQueue(12)
	//queue.show()
	_, _ = queue.deQueue()
	_, _ = queue.deQueue()
	_ = queue.enQueue(90)
	_ = queue.enQueue(56)
	fmt.Printf("%s\n", "==================")
	queue.show()

	_, _ = queue.deQueue()
	_, _ = queue.deQueue()
	fmt.Printf("%s\n", "==================")
	queue.show()
	_, _ = queue.deQueue()
	_, _ = queue.deQueue()
	_, _ = queue.deQueue()
	_, _ = queue.deQueue()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", "==================")
	queue.show()
}
