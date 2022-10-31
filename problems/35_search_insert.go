package problem

func searchInsert(nums []int, target int) int {
	var index int
	for i, num := range nums {
		if len(nums)-1 == i && num < target {
			index = i + 1
			break
		}

		if num == target {
			index = i
			break
		}

		if num > target {
			if i == 0 && num > target {
				index = 0
				break
			}

			if nums[i-1] == target {
				index = i - 1
				break
			}

			index = i
			break
		}
	}

	return index
}
