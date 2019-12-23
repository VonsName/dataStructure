package main

import "fmt"

func binarySearch(src []int, target int) (index int, times int) {
	low := 0
	high := len(src) - 1
	for low <= high {
		times++
		// mid := (low + high) / 2

		// 插值查找法  适用于所查找的集合中的元素分布均匀的时候
		mid := low + (((target - src[low]) / (src[high] - src[low])) * (high - low))
		if target < src[mid] {
			high = mid - 1
		} else if target > src[mid] {
			low = mid + 1
		} else {
			return mid, times
		}
	}
	fmt.Printf("times:=%d\n", times)
	return 0, times
}
func main() {

	a := []int{2, 5, 6, 24, 35, 47, 59, 62, 73, 88, 99}
	target := 99
	search, times := binarySearch(a, target)
	fmt.Printf("target:%d index=%d times=%d\n", target, search, times)
}
