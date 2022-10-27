package main

import "fmt"

func main() {
	var a = []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Unsorted: ", a)
	selectionSort(a)
	fmt.Println("Sorted: ", a)
}

func selectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
