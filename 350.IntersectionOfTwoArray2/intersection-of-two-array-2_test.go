package IntersectionOfTwoArray2

import (
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	num1 := []int{1,2,2,1}
	num2 := []int{2,2}
	res := intersect(num1, num2)
	if reflect.DeepEqual(res, []int{2,2}) {
		t.Errorf("failed")
	}
}
