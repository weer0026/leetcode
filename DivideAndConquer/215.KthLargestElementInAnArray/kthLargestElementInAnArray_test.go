package KthLargestElementInAnArray

import "testing"

func TestOutput(t *testing.T) {
	nums := []int{2,1}
	k := 2
	re := findKthLargest(nums, k)

	if re != 1 {
		t.Errorf("output: %v, fail", re)
	}
}
