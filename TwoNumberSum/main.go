package main

import "fmt"

func main()  {
	nums := []int{2, 7, 11, 15}

	a := twoSum1(nums, 9)
	b := twoSum2(nums, 9)

	fmt.Println(a, b)
}

// Solution 1 
func twoSum1(nums []int, target int) []int {
	numsMap := map[int]int{}
	for i, num := range nums {
		potentialMatch := target - num
		if j, found := numsMap[potentialMatch]; found {
			return []int{i, j}
		}
		numsMap[num] = i
	}
	return []int{}
}

// Solution 2
func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		firstNumber := nums[i]
		for j := i + 1; j < len(nums); j++ {
			secondNumber := nums[j]
			if firstNumber + secondNumber == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}