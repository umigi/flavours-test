package main

import "fmt"

func main() {

	var arr = []string{"abcdde", "baccd", "eeabg"}
	var lenArr int
	var mapChar = map[string]map[int]int{}
	var result = []string{}

	fmt.Println("Array", arr)
	lenArr = len(arr)

	for index, word := range arr {
		for _, char := range word {
			if mapChar[string(char)] == nil {
				mapChar[string(char)] = map[int]int{}
			}
			mapChar[string(char)][index]++
		}
	}

	for charIndex, charAppereance := range mapChar {
		if len(charAppereance) == lenArr {
			result = append(result, charIndex)
		}

	}
	fmt.Println("Words that appear in every array :", result)
	fmt.Println("Output :", len(result))
}
