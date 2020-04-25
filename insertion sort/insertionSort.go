package main

import "fmt"

func main() {
	arr := []int{8,5,6,9,2,4}

	fmt.Println(insertionSort(arr))
}

func insertionSort(arr []int) []int {
	for i:=1 ; i < len(arr) ; i++ {
		var compElement = arr[i]
		var j = i -1

		for j >= 0 && compElement < arr[j] {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = compElement
	}

	return arr
}