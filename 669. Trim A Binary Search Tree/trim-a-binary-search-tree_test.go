package trim_a_binary_search_tree

import (
	"fmt"
	"testing"
)

func TestOutput(t *testing.T) {
	var root = new(TreeNode)
	root.Val = 1
	var left = new(TreeNode)
	left.Val = 0
	var right = new(TreeNode)
	right.Val = 2
	root.Left = left
	root.Right = right

	ret := trimBST(root, 1, 2)
	fmt.Println(ret)
}
