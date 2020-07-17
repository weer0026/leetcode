package SearchInsertPosition

import "testing"

func TestOutput(t *testing.T) {
	input := []int{1, 3, 5, 6}
	target := 5
	re := searchInsert(input, target)
	if re != 2 {
		t.Errorf("failed! %v", re)
	}
}

func TestOutput1(t *testing.T) {
	input := []int{1, 3, 5, 6}
	target := 2
	re := searchInsert(input, target)
	if re != 1 {
		t.Errorf("failed! %v", re)
	}
}

func TestOutput2(t *testing.T) {
	input := []int{1, 3, 5, 6}
	target := 7
	re := searchInsert(input, target)

	if re != 4 {
		t.Errorf("failed! %v", re)
	}
}

func TestOutput3(t *testing.T) {
	input := []int{1, 3, 5, 6}
	target := 0
	re := searchInsert(input, target)
	if re != 0 {
		t.Errorf("failed!")
	}
}
