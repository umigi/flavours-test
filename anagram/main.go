package main

import "fmt"

func main() {
	var a string = "anagram"
	var b string = "nagaram"
	var listChar = []string{}
	var countCharA = map[string]int{}
	var countCharB = map[string]int{}

	fmt.Println(a)
	fmt.Println(b)

	if len(a) != len(b) {
		fmt.Println("False")
		return
	}

	for _, char := range a {
		listChar = append(listChar, string(char))
		countCharA[string(char)]++
	}

	for _, char := range b {
		countCharB[string(char)]++
	}

	for _, char := range listChar {
		if countCharA[char] != countCharB[char] {
			fmt.Println("False")
			return
		}
	}

	fmt.Println("True")
	return

}
