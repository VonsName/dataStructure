package main

import "fmt"

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
func kmp(fullStr string, subStr string) {

}
func testMatchSubString() {
	fullStr := "goodgoogle33"
	subStr := "google"
	success := matchSubString(fullStr, subStr)
	fmt.Printf("match=%v\n", success)
}
func main() {
	testMatchSubString()
}
