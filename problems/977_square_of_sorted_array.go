package problem

func sortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}
