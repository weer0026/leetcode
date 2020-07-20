package TwoSum2

import (
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	re := twoSum(input, target)

	if !reflect.DeepEqual(re, []int{1, 2}) {
		t.Errorf("failed %v", re)
	}
}
