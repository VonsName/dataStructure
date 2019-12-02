package main

type SqList struct {
	data []int
	size int
}

func newSqList(initialize int) *SqList {

	return &SqList{
		data: make([]int, initialize),
		size: 0,
	}
}

func main() {

}
