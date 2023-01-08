package main

import "fmt"

func main() {

	var t int = 17
	var arr = []int{2, 7, 11, 15}

	fmt.Println("Sum :", t)
	fmt.Println("Array :", arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {

			if arr[i]+arr[j] == t {
				fmt.Printf("Result : [%d, %d]", i, j)
				return
			}
		}
	}

}
