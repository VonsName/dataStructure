package main

import "fmt"

func main() {
	directSort("desc")
}

/**
直接插入排序
*/
func directSort(mode string) {
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
