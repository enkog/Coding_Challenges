package main

import (
	"fmt"
	"sort"
)

func main()  {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(SortedSquaredArray(array))
}

func SortedSquaredArray(array []int) []int {
	var sortedSquaresArray = make([]int, len(array)) 
	
	for i, el := range array {
		sortedSquaresArray[i] = el * el
	}

	sort.Ints(sortedSquaresArray)
	return sortedSquaresArray
}