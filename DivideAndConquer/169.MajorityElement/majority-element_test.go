package MajorityElement

import "testing"

func TestOutput(t *testing.T) {
	nums := []int{3, 2, 3}
	re := majorityElement(nums)
	if re != 3 {
		t.Errorf("output : %v, fail!", re)
	}

	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	re2 := majorityElement(nums2)
	if re2 != 2 {
		t.Errorf("output : %v, fail!", re)
	}

	nums3 := []int{6, 5, 5}
	re3 := majorityElement(nums3)
	if re3 != 5 {
		t.Errorf("output : %v, fail!", re)
	} else {
		t.Logf("output : %v", re)
	}
}

func TestOutputBM(t *testing.T) {
	nums := []int{3, 2, 3}
	re := majorityElementBM(nums)
	if re != 3 {
		t.Errorf("output : %v, fail!", re)
	}

	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	re2 := majorityElementBM(nums2)
	if re2 != 2 {
		t.Errorf("output : %v, fail!", re)
	}

	nums3 := []int{6, 5, 5}
	re3 := majorityElementBM(nums3)
	if re3 != 5 {
		t.Errorf("output : %v, fail!", re)
	} else {
		t.Logf("output : %v", re)
	}
}
