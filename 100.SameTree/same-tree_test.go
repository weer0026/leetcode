package SameTree

import (
	"testing"
)

func TestOutput(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	p := sortedArrayToBST(a)
	q := sortedArrayToBST(b)
	res := isSameTree(p, q)

	if res != true {
		t.Errorf("failed!")
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	return processNode(nums, 0, len(nums) - 1)
}

func processNode(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// 选取中间右侧的数为根节点
	// 整除算法
	mid := (left + right + 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = processNode(nums, left, mid - 1)
	root.Right = processNode(nums, mid + 1, right)
	return root
}
