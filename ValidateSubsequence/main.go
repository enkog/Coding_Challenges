package main

import "fmt"

func main() {
  arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	seq := []int{1, 6, -1, 10}

	fmt.Println(IsValidSubsequence1(arr, seq))
	fmt.Println(IsValidSubsequence2(arr, seq))
}

// Solution 1 
func IsValidSubsequence1(array []int, sequence []int) bool {
  arrayIndex := 0
	sequenceIndex := 0

	for arrayIndex < len(array) && sequenceIndex < len(sequence) {
		if array[arrayIndex] == sequence[sequenceIndex] {
			sequenceIndex += 1
		}
		arrayIndex += 1
	}
	return sequenceIndex == len(sequence)
}

// Solution 2
func IsValidSubsequence2(array []int, sequence []int) bool {
	sequenceIndex := 0
	for _, el := range array {
		if el == sequence[sequenceIndex] {
			sequenceIndex += 1
		}		
		if sequenceIndex == len(sequence) {
			return true
		}
	}
	return false
}