package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	RightDepth := maxDepth(root.Right)
	return max(leftDepth, RightDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
