package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	size  int
	cap   int
	front int
	back  int
	data  []int
}

func NewCircleQueue(initCap int) *CircleQueue {

	queue := &CircleQueue{
		size:  0,
		cap:   initCap,
		front: 0,
		back:  0,
		data:  nil,
	}
	return queue
}

func (cq *CircleQueue) enQueue(d int) (err error) {
	if (cq.back+1)%cq.cap == cq.front {
		return fmt.Errorf("%s", errors.New("当前队列已满"))
	}
	cq.data = append(cq.data, d)
	cq.back = (cq.back + 1) % cq.cap
	cq.size++
	return nil
}

func (cq *CircleQueue) deQueue() (int, error) {
	if cq.size == 0 {
		return 0, fmt.Errorf("%s", errors.New("当前队列已空"))
	}
	d := cq.data[cq.front]
	cq.front = (cq.front + 1) % cq.cap
	cq.size--
	return d, nil
}

func main() {

	queue := NewCircleQueue(5)
	_ = queue.enQueue(1)
	_ = queue.enQueue(2)
	_ = queue.enQueue(3)
	_ = queue.enQueue(4)
	_ = queue.enQueue(5)

	a, e := queue.deQueue()
	b, e := queue.deQueue()
	if e != nil {
		panic(e)
	}
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	_ = queue.enQueue(7)
	_ = queue.enQueue(8)
	// e = queue.enQueue(9)
	// if e != nil {
	// 	panic(e)
	// }

	c, e := queue.deQueue()
	d, e := queue.deQueue()
	f, e := queue.deQueue()
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
	fmt.Printf("%d\n", f)

	d, _ = queue.deQueue()
	f, _ = queue.deQueue()

	fmt.Printf("%d\n", d)
	fmt.Printf("%d\n", f)

}
