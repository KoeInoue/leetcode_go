package problem

func twoSum(numbers []int, target int) []int {
	for index, num := range numbers {
		for i, newNumber := range numbers {
			if index == i {
				continue
			}
			if num+newNumber == target {
				return []int{index, i}
			}
		}
	}

	return nil
}
