package main

import "fmt"

func main() {
	// directInsertSort("desc")

	a := []int{6, 5, 2, 13, 6, 23, 3, 8, 9}
	bubbleSort1(a)
	reverse(a)
	fmt.Printf("%v\n", a)
	// bothWayBubbleSort(a)
	// fmt.Printf("%v\n", a)
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
