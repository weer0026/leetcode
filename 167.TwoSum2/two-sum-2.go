package TwoSum2

// 双指针
func twoSum(numbers []int, target int) []int {
	if numbers == nil {
		return nil
	}
	low, high := 0, len(numbers)-1
	for low < high {
		tmp := numbers[low] + numbers[high]
		if tmp == target {
			return []int{low + 1, high + 1}
		} else if tmp < target {
			low++
		} else {
			high--
		}
	}
	return []int{low + 1, high + 1}
}
