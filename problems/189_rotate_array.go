package problem

import "fmt"

func rotate(nums []int, k int) {
	if k < 1 || len(nums) == 1 || k == len(nums) {
		fmt.Printf("nums: %v \n", nums)
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}
	// {1,2,3,4,5,6,7} k=3
	// Reverse the whole array
	reverseArray(nums, 0, len(nums)-1)
	// {7,6,5,4,3,2,1}
	// Reverse the first k elements
	reverseArray(nums, 0, k-1)
	// {5,6,7,4,3,2,1}
	// Reverse the rest of the array
	reverseArray(nums, k, len(nums)-1)
	// {5,6,7,1,2,3,4}

	fmt.Printf("nums: %v \n", nums)
}

func reverseArray(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}
}
