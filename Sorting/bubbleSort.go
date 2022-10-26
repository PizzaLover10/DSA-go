package main

import "fmt"

func main() {
	var a = []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Unsorted: ", a)
	bubbleSort(a)
	fmt.Println("Sorted: ", a)
}

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
