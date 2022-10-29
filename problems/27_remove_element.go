package problem

func removeElement(nums []int, val int) int {
	a := 0

	if len(nums) == 0 {
		return a
	}

	for b := 0; b < len(nums); b++ {
		if nums[a] == val {
			nums = append(nums[:a], nums[a+1:]...)
			if len(nums) > a {
				nums = append(nums, nums[a])
			}
		} else {
			a++
		}
	}

	return a
}
