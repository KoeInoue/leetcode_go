package problem

func search(nums []int, target int) int {
	result := -1

	for i, num := range nums {
		if num == target {
			result = i
			break
		}
	}

	return result
}
