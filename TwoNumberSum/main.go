func twoSum(nums []int, target int) []int {
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