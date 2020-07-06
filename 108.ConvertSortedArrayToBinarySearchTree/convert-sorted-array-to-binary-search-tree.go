package ConvertSortedArrayToBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
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