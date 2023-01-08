package main

import (
	"fmt"
	"strings"
)

func findUniqueWord(sentence string) []string {

	var unique = []string{}

	for _, word := range strings.Split(sentence, " ") {

		skip := false
		for _, uniqueWord := range unique {
			if word == uniqueWord {
				skip = true
				break
			}
		}

		if skip == false {
			unique = append(unique, word)
		}
	}

	return unique
}

func findUniqueChar(sentence string) []string {

	var unique = []string{}

	for _, word := range sentence {

		skip := false
		for _, uniqueWord := range unique {
			if string(word) == uniqueWord {
				skip = true
				break
			}
		}

		if skip == false {
			unique = append(unique, string(word))
		}
	}

	return unique
}

func main() {

	var sentence string = "hello world world hello"
	var pattern string = "abba"
	var uniqueWords = []string{}
	var uniqueChar = []string{}

	uniqueWords = findUniqueWord(sentence)
	uniqueChar = findUniqueChar(pattern)

	if len(uniqueChar) != len(uniqueWords) {
		fmt.Println("False")
		return
	}

	var wordPairing = map[string]string{}
	var generatedSentence = ""

	for index, char := range uniqueChar {
		wordPairing[char] = uniqueWords[index]
	}

	for index, char := range pattern {
		if index == 0 {
			generatedSentence = wordPairing[string(char)]
		} else {
			generatedSentence += " " + wordPairing[string(char)]
		}

	}

	fmt.Println("Pattern :", pattern)
	fmt.Println("Sentence :", sentence)
	fmt.Println("Generated Sentence :", generatedSentence)

	if generatedSentence == sentence {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
