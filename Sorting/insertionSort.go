package main

import "fmt"

func main() {
	var a = []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Unsorted: ", a)
	insertionSort(a)
	fmt.Println("Sorted: ", a)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
