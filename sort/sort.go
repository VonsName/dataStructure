package main

import (
	"fmt"
)

func main() {
	// directInsertSort("desc")

	// a := []int{6, 5, 2, 13, 6, 23, 3, 8, 9}
	// s := a[:]
	// // bubbleSort1(a)
	// // reverse(a)
	// // fastSort(a, 0, len(a)-1)
	// // bothWayBubbleSort(a)
	// // fmt.Printf("%v\n", a)
	// // shellSort(a)
	// //
	// // printOut(345)
	// // var res []int
	// // fmt.Printf("%v ", a)
	// quickSort(a)
	// fmt.Printf("a=%v\n", a)
	// fmt.Printf("origin=%v\n", s)
	// s = append(s[:0], 1, 2, 3, 4, 5, 6, 7, 8, 9)
	// fmt.Printf("append1=%v \n", s)
	// s = append(s[:], 1, 2, 3, 4, 5, 6, 7, 8, 9)
	// fmt.Printf("append2=%v \n", s)
	a := []int{16, 7, 3, 20, 17, 8}
	heapSort(a)
	a[len(a)-1], a[0] = a[0], a[len(a)-1]
	fmt.Printf("%v\n", a)
}

func printOut(n int) {
	if n > 10 {
		printOut(n / 10)
	}
	fmt.Printf("%v ", n%10)
}
func shellSort(a []int) {
	var j int
	for gap := len(a) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(a); i++ {
			temp := a[i]
			for j = i; j >= gap && temp < a[j-gap]; j -= gap {
				a[j] = a[j-gap]
			}
			a[j] = temp
		}
	}

	fmt.Printf("%v\n", a)
}

/**
分区
*/
func partition(a []int, i int, j int) int {
	r := a[i]
	for i < j {
		for j > i && a[j] > r {
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}
		for i < j && a[i] < r {
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	a[i] = r
	fmt.Printf("%v\n", a)
	return i
}

/**
快速排序
*/
func fastSort(a []int, low int, high int) {
	if low < high {
		i2 := partition(a, low, high)
		fastSort(a, low, i2-1)
		fastSort(a, i2+1, high)
	}
}

func reverse(src []int) {
	n := len(src) / 2
	j := len(src)
	for i := 0; i < n; i++ {
		temp := src[i]
		src[i] = src[j-1-i]
		src[j-1-i] = temp
	}
}

/**
双向冒泡
*/
func bothWayBubbleSort(a []int) {
	swap := true
	i, j, k := 0, 0, 0
	for swap {
		swap = false

		//
		for k = len(a) - i - 1; k > i; k-- {
			if a[k] < a[k-1] {
				temp := a[k]
				a[k] = a[k-1]
				a[k-1] = temp
				swap = true
				fmt.Printf("%s\n", "=====================")
			}
		}

		for j = i + 1; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				swap = true
				fmt.Printf("%s\n", "=====================")
			}
		}
		i++
	}
}

/**
冒泡
*/
func bubbleSort1(a []int) {

	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				temp := a[j]
				a[j] = a[j-1]
				a[j-1] = temp
			}
		}
	}
	fmt.Printf("%v\n", a)
}

/**
冒泡
*/
func bubbleSort(a []int) {

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] < a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
			// fmt.Printf("%v\n", a)
		}
	}

	fmt.Printf("%v\n", a)
}

func quickSort(items []int) {
	if len(items) > 1 {
		var smaller []int
		var larger []int
		var same []int

		chosen := items[len(items)/2]
		for i := 0; i < len(items); i++ {
			v := items[i]
			if v < chosen {
				smaller = append(smaller, v)
			} else if v > chosen {
				larger = append(larger, v)
			} else {
				same = append(same, v)
			}
		}
		quickSort(smaller)
		quickSort(larger)
		items = append(items[:0], smaller...)
		items = append(items, same...)
		items = append(items, larger...)
	}
}

/**
直接插入排序
*/
func directInsertSort(mode string) {
	b := make([]int, 1)
	a := []int{1, 5, 2, 13, 6, 23, 3, 8, 9}

	var j int
	for i := 1; i < len(a); i++ {
		switch mode {
		case "asc":
			if a[i] < a[i-1] {
				b[0] = a[i]
				for j = i - 1; j >= 0 && b[0] < a[j]; j-- { // 寻找合适的插入位置
					a[j+1] = a[j]
				}
				// 插入元素
				a[j+1] = b[0]
			}

		case "desc":
			if a[i] > a[i-1] {
				b[0] = a[i]
				for j = i - 1; j >= 0 && b[0] > a[j]; j-- {
					a[j+1] = a[j]
				}
				a[j+1] = b[0]
			}
		}
	}

	fmt.Printf("%v\n", a)
}

/**
堆排序 利用完全二叉树的性质
 将完全二叉树按照从顶之下,从左至右的顺序存入数组q
 i=0时为根节点,无双亲,否则qi的双亲为i-1/2
 qi的左孩子为2i+1,右孩子为2i+2
*/

func heapSort(a []int) {
	k := 0
	for i := len(a) - 1; i >= 0; i-- {
		buildHeap(a, i)
		a[i], a[k] = a[k], a[i]
	}
}

/**
构建堆
*/
func buildHeap(a []int, i int) {
	tempI := i
	for parentIndex := (i - 1) / 2; parentIndex >= 0 && i >= 0; parentIndex = (i - 1) / 2 {
		leftChildIndex := 2*parentIndex + 1
		rightChildIndex := 2*parentIndex + 2

		m := 0
		if a[leftChildIndex] > a[parentIndex] {
			m = leftChildIndex
			a[parentIndex], a[leftChildIndex] = a[leftChildIndex], a[parentIndex]
		}
		if rightChildIndex < len(a) {
			m = rightChildIndex
			if a[rightChildIndex] > a[parentIndex] {
				a[parentIndex], a[rightChildIndex] = a[rightChildIndex], a[parentIndex]
			}
		}
		if m != 0 {
			childSort(m, tempI, a)
		}
		i -= 2
	}
}

func childSort(parentIndex int, i int, a []int) {

	leftChildIndex := 2*parentIndex + 1
	rightChildIndex := 2*parentIndex + 2

	if leftChildIndex <= i {
		if a[leftChildIndex] > a[parentIndex] {
			a[parentIndex], a[leftChildIndex] = a[leftChildIndex], a[parentIndex]
		}
		childSort(leftChildIndex, i, a)
	}
	if rightChildIndex <= i {
		if a[rightChildIndex] > a[parentIndex] {
			a[parentIndex], a[rightChildIndex] = a[rightChildIndex], a[parentIndex]
		}
		childSort(rightChildIndex, i, a)
	}
}
