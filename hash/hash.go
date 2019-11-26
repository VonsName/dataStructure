package main

import "fmt"

func main() {
	str := "HelloWorld"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	b := []byte(str)
	println()
	for i := 0; i < len(b); i++ {
		fmt.Printf("%v ", b[i])
	}
}
