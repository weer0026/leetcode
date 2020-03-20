package Average_of_Levels_in_Binary_Tree

import (
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	root, l1, r1, l2_3, r2_4 := new(TreeNode), new(TreeNode), new(TreeNode), new(TreeNode), new(TreeNode)
	root.Val = 3
	l1.Val = 9
	root.Left = l1
	r1.Val = 20
	l2_3.Val, r2_4.Val = 15, 7
	r1.Left, r1.Right = l2_3, r2_4
	root.Right = r1
	arr := averageOfLevels(root)
	expected := []float{3, 14.5, 11}

	if !reflect.DeepEqual(arr, expected) {
		t.Error("unexpected result")
	}
}
