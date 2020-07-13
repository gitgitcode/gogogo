package main

import (
	"fmt"
)

func main() {
	toSorte := []int{34, 56, 7, 8, 43, 45, 22, 4444, 144, 9998}
	bubbleSort(toSorte)
	fmt.Println(toSorte)
}

func bubbleSort(arr []int) {

	for {
		sorted := true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		if sorted == true {
			break
		}
	}
}
