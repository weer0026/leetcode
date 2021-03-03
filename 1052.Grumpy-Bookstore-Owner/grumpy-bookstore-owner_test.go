package _052_Grumpy_Bookstore_Owner

import "testing"

// customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
func TestMaxSatisfied(t *testing.T) {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	x := 3
	excepted := 16
	output := maxSatisfied(customers, grumpy, x)
	if output != excepted {
		t.Errorf("结果不符合预期！exceted:%v, output:%v", excepted, output)
	}
}
