package _98_HouseRobber

import "testing"

func TestRob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	re := rob2(nums)
	if re != 4 {
		t.Errorf("failed! output: %v", re)
	}
}
