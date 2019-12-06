package main

import (
	"fmt"
)

func matchSubString(fullStr string, subStr string) (success bool) {
	fullMatch := 0
	i := 0
	j := 0
	for i < len(fullStr) && j < len(subStr) {
		if fullMatch == len(subStr) {
			return
		}
		if fullStr[i] == subStr[j] {
			i++
			j++
			fullMatch++
			success = true
		} else {
			i = i - (j - 1)
			j = 0
			fullMatch = 0
			success = false
		}
	}
	return
}
func kmp(fullStr string, subStr string) (int, int) {
	pm := pmt(subStr)
	i, j := 0, 0
	for i < len(fullStr) && j < len(subStr) {
		if fullStr[i] == subStr[j] {
			i++
			j++
		} else {
			if j != 0 {
				j = pm[j-1]

				// 只是为了打印最长前缀和后缀
				// temp := pm[j-1]
				// if temp != 0 {
				// 	t := j - 1
				// 	stack := linear_list.NewLinkedStack()
				// 	for k := temp; k > 0; k-- {
				// 		stack.Push(string(subStr[t]))
				// 		t--
				// 	}
				// 	stack.Show()
				// }
				// 只是为了打印最长前缀和后缀

			} else {
				i++
			}
		}
	}
	if j == len(subStr) {
		return i - j, i - 1
	}
	return -1, -1
}

/**
构建Pmt数组
*/
func pmt(subStr string) []int {
	pm := make([]int, len(subStr))
	i, j := 0, 1
	pm[i] = 0
	for j < len(subStr) {
		if subStr[i] == subStr[j] {
			i++
			pm[j] = i // pm[j]=i+1
			j++
		} else {
			if i != 0 {
				i = pm[i-1]
			} else {
				pm[j] = 0
				j++
			}
		}
	}
	return pm
}

func testKmp() {
	text := "abaadbabaabcacdc"
	pattern := "abaabcac"
	startIndex, endIndex := kmp(text, pattern)
	fmt.Printf("substr:%s at fullstr:%s startIndex=%d endIndex=%d\n", pattern, text, startIndex, endIndex)
}
func testPmt() {
	// str := "aabaabaaa"
	// str := "abcdabca"
	str := "abaabcac"
	ints := pmt(str)
	fmt.Printf("%v\n", ints)
}
func testMatchSubString() {
	fullStr := "goodgoogle33"
	subStr := "google"
	success := matchSubString(fullStr, subStr)
	fmt.Printf("match=%v\n", success)
}
func main() {
	// testMatchSubString()
	testPmt()
	testKmp()
}
