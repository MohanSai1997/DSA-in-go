package main

import (
	"fmt"
)

func main() {
	var arr = []int{6, 5, 12, 10, 9, 1}

	fmt.Println(mergeSort(arr))
}

func mergeSort(arr []int) []int{
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr)/2
	
	return merge( mergeSort(arr[:mid]), mergeSort(arr[mid:]) )
}

func merge(left []int, right []int) []int {
	sliceSize, i, j := len(left) + len(right), 0, 0
	slice := make([]int, sliceSize, sliceSize)

	for k:=0; k < sliceSize ; k++ {

		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}

	return slice
}