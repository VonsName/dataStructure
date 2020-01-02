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
	a := []int{35, 20, 55, 60, 40, 38, 10, 70, 40}
	heapSort(a)
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
	visit := make([]bool, len(a))
	for i := 0; i < len(visit); i++ {
		visit[i] = false
	}
	for i := len(a) - 1; i >= 0; i-- {
		buildHeap(a, i, visit)
		visit[i] = true
		// 当父节点是0的时候 直接比较父节点以及左右孩子即可
		if i == 2 || i == 1 {
			break
		}
		a[k], a[i] = a[i], a[k]
	}
}

/**
构建堆
*/
func buildHeap(a []int, i int, visit []bool) {
	temp := i
	parentIndex := (i - 1) / 2
	if parentIndex == 0 {
		leftChildIndex := 2*parentIndex + 1
		rightChildIndex := 2*parentIndex + 2
		if a[leftChildIndex] < a[parentIndex] {
			swap(parentIndex, leftChildIndex, a)
		}
		if a[rightChildIndex] < a[parentIndex] {
			swap(parentIndex, rightChildIndex, a)
		}
		if a[rightChildIndex] < a[leftChildIndex] {
			swap(leftChildIndex, rightChildIndex, a)
		}
	} else {
		for ; parentIndex >= 0 && i >= 0; parentIndex = (i - 1) / 2 {
			leftChildIndex := 2*parentIndex + 1
			rightChildIndex := 2*parentIndex + 2
			if a[leftChildIndex] > a[parentIndex] && !visit[leftChildIndex] {
				swap(parentIndex, leftChildIndex, a)
				childSort(leftChildIndex, temp, a, visit)
			}
			if (rightChildIndex) < len(a) && !visit[rightChildIndex] {
				if a[rightChildIndex] > a[parentIndex] {
					swap(parentIndex, rightChildIndex, a)
					childSort(leftChildIndex, temp, a, visit)
				}
			}
			i -= 2
		}
	}
}

func childSort(parentIndex int, i int, a []int, visit []bool) {

	leftChildIndex := 2*parentIndex + 1
	rightChildIndex := 2*parentIndex + 2

	if leftChildIndex < len(a) && !visit[leftChildIndex] {
		if a[leftChildIndex] > a[parentIndex] {
			swap(parentIndex, leftChildIndex, a)
		}
		childSort(leftChildIndex, i, a, visit)
	}
	if rightChildIndex < len(a) && !visit[rightChildIndex] {
		if a[rightChildIndex] > a[parentIndex] {
			swap(parentIndex, rightChildIndex, a)
		}
		childSort(rightChildIndex, i, a, visit)
	}
}

func swap(i int, j int, a []int) {
	a[i], a[j] = a[j], a[i]
}
