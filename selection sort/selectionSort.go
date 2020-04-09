package main

import (
	"fmt"
)

func main() {
	var arrLength int
	fmt.Println("Enter the size of array: ")
	fmt.Scan(&arrLength)

	fmt.Println("Enter the array elemnent: ")
	arr := make([]int, arrLength, arrLength)
	for i:=0; i<arrLength; i++ {
		fmt.Scan(&arr[i])
	}

	var minimumIndex int

	for i:=0 ; i<arrLength ; i++ {
		minimumIndex = i

		for j:=i+1 ; j < arrLength ; j++ {
			if arr[j]>arr[minimumIndex] {
				minimumIndex = j
			}
		}
		arr[i], arr[minimumIndex] = arr[minimumIndex], arr[i]
	}

	fmt.Println(arr)
}
