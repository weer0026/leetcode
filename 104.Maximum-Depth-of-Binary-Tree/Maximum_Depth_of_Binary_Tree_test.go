package Maximum_Depth_of_Binary_Tree

import (
	"testing"
)

func TestOutput(t *testing.T) {
	root := new(TreeNode)
	root.Val = 3
	l1 := new(TreeNode)
	l1.Val = 9
	root.Left = l1
	r1 := new(TreeNode)
	r1.Val = 20
	r2 := new(TreeNode)
	r2.Val = 15
	r3 := new(TreeNode)
	r3.Val = 7
	r1.Left = r2
	r1.Right = r3
	root.Right = r1
	depth := maxDepth(root)

	if depth != 3 {
		t.Error("fail")
	}
}
