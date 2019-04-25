package Add_Two_Numbers

import (
	"fmt"
	"testing"
)

func TestOutput(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 2
	l1_2 := new(ListNode)
	l1_2.Val = 4
	l1_3 := new(ListNode)
	l1_3.Val = 3
	l1_2.Next = l1_3
	l1.Next = l1_2

	l2 := new(ListNode)
	l2.Val = 5
	l2_2 := new(ListNode)
	l2_2.Val = 6
	l2_3 := new(ListNode)
	l2_3.Val = 4
	l2_2.Next = l2_3
	l2.Next = l2_2

	ret := addTwoNumbers(l1, l2)

	fmt.Printf("%v", ret)
}
