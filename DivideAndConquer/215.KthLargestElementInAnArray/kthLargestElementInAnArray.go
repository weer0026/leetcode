package KthLargestElementInAnArray

import (
	"fmt"
	"sort"
)

type arrayInterface []int

func (arr arrayInterface) Len() int {
	return len(arr)
}

func (arr arrayInterface) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr arrayInterface) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

/**
	sort and unique
 */
func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Sort(arrayInterface(nums))
	count := 0
	fmt.Printf("%v", nums)

	for i := len(nums) - 1; i >= 0; i-- {
		count++
		if count == k {
			return nums[i]
		}
	}
	return 0
}
