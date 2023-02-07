package main

import (
	"fmt"
)

func getParent(i int) int {
	return (i - 1) / 2
}

func getLeft(i int) int {
	return 2*i + 1
}

func getRight(i int) int {
	return 2 * (i + 1)
}

func swap(a int, b int) (int, int) {
	return b, a
}

func maxHeapify(a []int, i int) {
	heapSize := len(a)
	l := getLeft(i)
	r := getRight(i)
	var largest int
	if l < heapSize && a[l] > a[i] {
		largest = l
	} else {
		largest = i
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	fmt.Println(a[largest])
	if largest != i {
		a[i], a[largest] = swap(a[i], a[largest])
		maxHeapify(a, largest)
	}
}

func main() {
	fmt.Println("Heap and different operations")
	a := []int{1, 16, 14, 10, 7, 8, 9, 4, 2, 3}
	fmt.Println(a)
	maxHeapify(a, 0)
	fmt.Println("after max heapify", a)

}
