package trim_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if nil == root {
		return root
	}
	val := root.Val

	if val >= L && val <= R {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}

	temp := root

	if val < L {
		temp = root.Right
	} else if val > R {
		temp = root.Left
	}
	return trimBST(temp, L, R )
}
