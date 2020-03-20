package MaximumSubarray

import "testing"

func TestOutput(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	re := maxSubArray(arr)

	if re != 6 {
		t.Error("fail")
	}
}
