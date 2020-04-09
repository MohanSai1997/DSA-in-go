package main

import (
	"fmt"
)

func main() {
	var arrLength int
	fmt.Println("Enter the Array Length: ")
	fmt.Scan(&arrLength)

	var arr = make([]int, arrLength, arrLength)
	fmt.Println("Enter the Array: ")

	for i := 0 ; i<arrLength ; i++{
		fmt.Scan(&arr[i])
	}

	for i:=0 ; i<arrLength ; i++{
		for j:=0; j<arrLength - i - 1; j++{
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Bubbled array: ", arr)
}
